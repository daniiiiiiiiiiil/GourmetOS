package server

import (
	"GourmetOS/internal/handler"
	"GourmetOS/internal/middleware"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HTTPServer struct {
	pool                *pgxpool.Pool
	orderHandler        *handler.OrderHandler
	dishHandler         *handler.DishHandler
	tableHandler        *handler.TableHandler
	paymentHandler      *handler.PaymentHandler
	kitchenHandler      *handler.KitchenHandler
	notificationHandler *handler.NotificationHandler
}

func NewHTTPServer(
	pool *pgxpool.Pool,
	orderHandler *handler.OrderHandler,
	dishHandler *handler.DishHandler,
	tableHandler *handler.TableHandler,
	paymentHandler *handler.PaymentHandler,
	kitchenHandler *handler.KitchenHandler,
	notificationHandler *handler.NotificationHandler,
) *HTTPServer {
	return &HTTPServer{
		pool:                pool,
		orderHandler:        orderHandler,
		dishHandler:         dishHandler,
		tableHandler:        tableHandler,
		paymentHandler:      paymentHandler,
		kitchenHandler:      kitchenHandler,
		notificationHandler: notificationHandler,
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (server *HTTPServer) Start() error {
	router := mux.NewRouter()

	router.Use(middleware.DBConnection(server.pool))

	router.Use(corsMiddleware)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/health", server.healthCheck).Methods("GET")

	api.HandleFunc("/orders", server.orderHandler.CreateOrder).Methods("POST")
	api.HandleFunc("/orders", server.orderHandler.GetAllOrders).Methods("GET")
	api.HandleFunc("/orders/active", server.orderHandler.GetActiveOrders).Methods("GET")
	api.HandleFunc("/orders/{id}", server.orderHandler.GetOrder).Methods("GET")
	api.HandleFunc("/orders/{id}/history", server.orderHandler.GetOrderHistory).Methods("GET")
	api.HandleFunc("/orders/{id}/dishes", server.orderHandler.AddDish).Methods("POST")
	api.HandleFunc("/orders/{id}/dishes", server.orderHandler.RemoveDish).Methods("DELETE")
	api.HandleFunc("/orders/{id}/submit", server.orderHandler.SubmitToKitchen).Methods("PUT")
	api.HandleFunc("/orders/{id}/ready", server.orderHandler.MarkAsReady).Methods("PUT")
	api.HandleFunc("/orders/{id}/serve", server.orderHandler.ServeToTable).Methods("PUT")
	api.HandleFunc("/orders/{id}/pay", server.orderHandler.ProcessPayment).Methods("POST")
	api.HandleFunc("/orders/{id}/cancel", server.orderHandler.CancelOrder).Methods("PUT")
	api.HandleFunc("/orders/{id}/undo", server.orderHandler.UndoLastAction).Methods("POST")
	api.HandleFunc("/orders/{id}/redo", server.orderHandler.RedoLastAction).Methods("POST")
	api.HandleFunc("/orders/{id}/total", server.orderHandler.GetOrderTotal).Methods("GET")

	api.HandleFunc("/dishes", server.dishHandler.CreateDish).Methods("POST")
	api.HandleFunc("/dishes", server.dishHandler.GetAllDishes).Methods("GET")
	api.HandleFunc("/dishes/available", server.dishHandler.GetAvailableDishes).Methods("GET")
	api.HandleFunc("/dishes/vegetarian", server.dishHandler.GetVegetarianDishes).Methods("GET")
	api.HandleFunc("/dishes/vegan", server.dishHandler.GetVeganDishes).Methods("GET")
	api.HandleFunc("/dishes/gluten-free", server.dishHandler.GetGlutenFreeDishes).Methods("GET")
	api.HandleFunc("/dishes/search", server.dishHandler.SearchDishes).Methods("GET")
	api.HandleFunc("/dishes/category/{category}", server.dishHandler.GetDishesByCategory).Methods("GET")
	api.HandleFunc("/dishes/cuisine/{cuisine}", server.dishHandler.GetDishesByCuisine).Methods("GET")
	api.HandleFunc("/dishes/price", server.dishHandler.GetDishesByPriceRange).Methods("GET")
	api.HandleFunc("/dishes/{id}", server.dishHandler.GetDish).Methods("GET")
	api.HandleFunc("/dishes/{id}", server.dishHandler.UpdateDish).Methods("PUT")
	api.HandleFunc("/dishes/{id}", server.dishHandler.DeleteDish).Methods("DELETE")
	api.HandleFunc("/dishes/{id}/availability", server.dishHandler.UpdateAvailability).Methods("PATCH")
	api.HandleFunc("/dishes/{id}/decorators", server.dishHandler.AddDecorator).Methods("POST")
	api.HandleFunc("/dishes/{id}/decorators", server.dishHandler.RemoveDecorator).Methods("DELETE")

	api.HandleFunc("/menu/tree", server.dishHandler.GetMenuTree).Methods("GET")
	api.HandleFunc("/menu/combo/{type}", server.dishHandler.GetComboMeal).Methods("GET")

	api.HandleFunc("/tables", server.tableHandler.CreateTable).Methods("POST")
	api.HandleFunc("/tables", server.tableHandler.GetAllTables).Methods("GET")
	api.HandleFunc("/tables/free", server.tableHandler.GetFreeTables).Methods("GET")
	api.HandleFunc("/tables/occupied", server.tableHandler.GetOccupiedTables).Methods("GET")
	api.HandleFunc("/tables/location/{location}", server.tableHandler.GetTablesByLocation).Methods("GET")
	api.HandleFunc("/tables/capacity/{capacity}", server.tableHandler.GetTablesByCapacity).Methods("GET")
	api.HandleFunc("/tables/{id}", server.tableHandler.GetTable).Methods("GET")
	api.HandleFunc("/tables/{id}", server.tableHandler.UpdateTable).Methods("PUT")
	api.HandleFunc("/tables/{id}", server.tableHandler.DeleteTable).Methods("DELETE")
	api.HandleFunc("/tables/{id}/occupy", server.tableHandler.OccupyTable).Methods("PATCH")
	api.HandleFunc("/tables/{id}/free", server.tableHandler.FreeTable).Methods("PATCH")
	api.HandleFunc("/tables/{id}/reserve", server.tableHandler.ReserveTable).Methods("PATCH")
	api.HandleFunc("/tables/{id}/cancel-reserve", server.tableHandler.CancelReservation).Methods("PATCH")

	api.HandleFunc("/payments", server.paymentHandler.ProcessPayment).Methods("POST")
	api.HandleFunc("/payments/methods", server.paymentHandler.GetPaymentMethods).Methods("GET")
	api.HandleFunc("/payments/discount", server.paymentHandler.ApplyDiscount).Methods("POST")
	api.HandleFunc("/payments/calculate", server.paymentHandler.CalculateFinalAmount).Methods("POST")
	api.HandleFunc("/payments/{id}/status", server.paymentHandler.GetPaymentStatus).Methods("GET")
	api.HandleFunc("/payments/{id}/refund", server.paymentHandler.RefundPayment).Methods("POST")
	api.HandleFunc("/payments/order/{orderId}", server.paymentHandler.GetPaymentByOrder).Methods("GET")
	api.HandleFunc("/payments/transaction/{txnId}", server.paymentHandler.GetPaymentByTransaction).Methods("GET")

	api.HandleFunc("/kitchen/receive", server.kitchenHandler.ReceiveOrder).Methods("POST")
	api.HandleFunc("/kitchen/cook/{id}", server.kitchenHandler.StartCooking).Methods("PUT")
	api.HandleFunc("/kitchen/ready/{id}", server.kitchenHandler.MarkAsReady).Methods("PUT")
	api.HandleFunc("/kitchen/queue", server.kitchenHandler.GetCookingQueue).Methods("GET")
	api.HandleFunc("/kitchen/time/{dishId}", server.kitchenHandler.GetCookingTime).Methods("GET")
	api.HandleFunc("/kitchen/status", server.kitchenHandler.GetKitchenStatus).Methods("GET")

	api.HandleFunc("/notifications/subscribe", server.notificationHandler.Subscribe).Methods("POST")
	api.HandleFunc("/notifications/unsubscribe", server.notificationHandler.Unsubscribe).Methods("POST")
	api.HandleFunc("/notifications/subscribers", server.notificationHandler.GetSubscribers).Methods("GET")
	api.HandleFunc("/notifications/send", server.notificationHandler.SendNotification).Methods("POST")
	api.HandleFunc("/notifications/order-created", server.notificationHandler.NotifyOrderCreated).Methods("POST")
	api.HandleFunc("/notifications/order-ready", server.notificationHandler.NotifyOrderReady).Methods("POST")
	api.HandleFunc("/notifications/order-served", server.notificationHandler.NotifyOrderServed).Methods("POST")
	api.HandleFunc("/notifications/order-paid", server.notificationHandler.NotifyOrderPaid).Methods("POST")

	handler := corsMiddleware(router)

	port := ":8080"
	if err := http.ListenAndServe(port, handler); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}

func (server *HTTPServer) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "ok",
		"service": "GourmetOS",
		"version": "1.0.0",
	})
}
