package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"GourmetOS/configs"
	"GourmetOS/internal/handler"
	"GourmetOS/internal/patterns/command"
	"GourmetOS/internal/patterns/factory"
	"GourmetOS/internal/patterns/observer"
	"GourmetOS/internal/patterns/proxy"
	"GourmetOS/internal/patterns/singleton"
	"GourmetOS/internal/repository/postgres"
	"GourmetOS/internal/server"
	"GourmetOS/internal/service"
	"GourmetOS/pkg/logger"
)

func main() {
	cfg, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	log := logger.NewLogger(cfg.Logger.Level)
	log.Info("Starting GourmetOS...", "version", cfg.App.Version)

	db := singleton.GetDBConnection(cfg.Database.URL)
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("failed to ping database", "error", err)
	}
	log.Info("Database connected successfully")

	pool := db.GetPool()

	orderRepo := postgres.NewOrderRepository(pool)
	dishRepo := postgres.NewDishRepository(pool)
	tableRepo := postgres.NewTableRepository(pool)
	paymentRepo := postgres.NewPaymentRepository(pool)
	customerRepo := postgres.NewCustomerRepository(pool)
	employeeRepo := postgres.NewEmployeeRepository(pool)
	ingredientRepo := postgres.NewIngredientRepository(pool)

	log.Info("Repositories initialized")

	subject := observer.NewOrderSubject()

	kitchenDisplay := observer.NewKitchenDisplay("Основная кухня")
	waiterNotifier := observer.NewWaiterNotifier("Алексей")
	customerNotifier := observer.NewCustomerNotifier("+7-999-123-45-67")
	adminNotifier := observer.NewAdminNotifier("Администратор")

	subject.RegisterObserver(kitchenDisplay)
	subject.RegisterObserver(waiterNotifier)
	subject.RegisterObserver(customerNotifier)
	subject.RegisterObserver(adminNotifier)

	log.Info("Observers registered", "count", 4)

	invoker := command.NewCommandInvoker()

	factoryRegistry := factory.NewFactoryRegistry()
	factoryRegistry.Register("italian", factory.NewItalianKitchen())
	factoryRegistry.Register("japanese", factory.NewJapaneseKitchen())
	factoryRegistry.Register("mexican", factory.NewMexicanKitchen())

	log.Info("Factory registry initialized", "kitchens", factoryRegistry.GetAvailableKitchens())

	menuProxy := proxy.NewMenuProxy()

	orderService := service.NewOrderService(
		orderRepo,
		paymentRepo,
		tableRepo,
		dishRepo,
		customerRepo,
		employeeRepo,
		subject,
		invoker,
		log.Logger,
	)

	dishService := service.NewDishService(
		dishRepo,
		ingredientRepo,
		factoryRegistry,
		menuProxy,
		log.Logger,
	)

	tableService := service.NewTableService(
		tableRepo,
		log.Logger,
	)

	paymentService := service.NewPaymentService(
		paymentRepo,
		orderRepo,
		invoker,
		log.Logger,
	)

	kitchenService := service.NewKitchenService(
		orderRepo,
		dishRepo,
		subject,
		log.Logger,
	)

	notificationService := service.NewNotificationService(
		subject,
		log.Logger,
	)

	log.Info("Services initialized successfully")

	orderHandler := handler.NewOrderHandler(orderService)
	dishHandler := handler.NewDishHandler(dishService)
	tableHandler := handler.NewTableHandler(tableService)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	kitchenHandler := handler.NewKitchenHandler(kitchenService)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	log.Info("Handlers initialized successfully")

	srv := server.NewHTTPServer(
		pool,
		orderHandler,
		dishHandler,
		tableHandler,
		paymentHandler,
		kitchenHandler,
		notificationHandler,
	)

	go func() {
		log.Info("Starting HTTP server", "port", cfg.Server.Port)
		if err := srv.Start(); err != nil {
			log.Fatal("Server failed", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = ctx

	log.Info("Server exited gracefully")
}
