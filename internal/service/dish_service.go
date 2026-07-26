package service

import (
	"GourmetOS/internal/domain"
	"GourmetOS/internal/patterns/abstract_factory"
	"GourmetOS/internal/patterns/composite"
	"GourmetOS/internal/patterns/decorator"
	"GourmetOS/internal/patterns/factory"
	"GourmetOS/internal/patterns/proxy"
	"GourmetOS/internal/repository/interfaceRepository"
	"GourmetOS/pkg/errors"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type DishService struct {
	dishService       interfaceRepository.DishRepository
	ingredientService interfaceRepository.IngredientRepository
	factoryRegistry   *factory.FactoryRegistry
	menuProxy         *proxy.MenuProxy
	logger            *zap.Logger
}

func NewDishService(
	dishService interfaceRepository.DishRepository,
	ingredientService interfaceRepository.IngredientRepository,
	factoryRegistry *factory.FactoryRegistry,
	menuProxy *proxy.MenuProxy,
	logger *zap.Logger) *DishService {
	return &DishService{
		dishService:       dishService,
		ingredientService: ingredientService,
		factoryRegistry:   factoryRegistry,
		menuProxy:         menuProxy,
		logger:            logger,
	}
}

func (s *DishService) CreateDish(ctx context.Context, conn *pgx.Conn, dish *domain.Dish) (*domain.Dish, error) {
	if err := dish.Validate(); err != nil {
		return nil, errors.ValidationError{
			Field:   "dish",
			Message: err.Error(),
		}
	}
	existing, err := s.dishService.GetByNameDish(ctx, conn, dish.Name, 1, 0)
	if err == nil && len(existing) > 0 {
		return nil, errors.BusinessError{
			Code:    "ErrDishAlreadyExists",
			Message: fmt.Sprintf("Блюдо с названием '%s' уже существует", dish.Name),
		}
	}

	if dish.CookingTime == 0 {
		kitchenFactory, err := s.factoryRegistry.GetFactory(dish.Cuisine)
		if err != nil {
			s.logger.Warn("cuisine not found, using default italian kitchen",
				zap.String("cuisine", dish.Cuisine))
			kitchenFactory, _ = s.factoryRegistry.GetFactory("italian")
		}

		var factoryDish factory.Dish

		switch dish.Category {
		case "pizza":
			factoryDish = kitchenFactory.CreatePizza()
		case "pasta":
			factoryDish = kitchenFactory.CreatePasta()
		case "salad":
			factoryDish = kitchenFactory.CreateSalad()
		case "drink":
			factoryDish = kitchenFactory.CreateDrink()
		}

		if factoryDish != nil && dish.CookingTime == 0 {
			dish.CookingTime = factoryDish.GetCookingTime()
		}
	}

	if dish.CookingTime == 0 {
		dish.CookingTime = 15
	}
	if !dish.IsAvailable && dish.IsAvailable == false {
		dish.IsAvailable = true
	}

	createdDish, err := s.dishService.CreateDish(ctx, conn, *dish)
	if err != nil {
		s.logger.Error("failed to create dish", zap.Error(err))
		return nil, errors.BusinessError{
			Code:    "ErrCreateDish",
			Message: "Не удалось создать блюдо: " + err.Error(),
		}
	}

	if s.menuProxy != nil {
		s.menuProxy.InvalidateCache()
	}

	return createdDish, nil
}

func (s *DishService) GetDish(ctx context.Context, conn *pgx.Conn, id int) (*domain.Dish, error) {
	if id <= 0 {
		return nil, errors.ValidationError{
			Field:   "id",
			Message: "ID не может быть меньше или равен 0",
		}
	}
	dish, err := s.dishService.GetByIDDish(ctx, conn, id)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetDish",
			Message: "Не удалось получить блюда по id" + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetAllDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	limitOffset(limit, offset)
	dish, err := s.dishService.GetAllDishes(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetAllDishes",
			Message: "Не удалось получить все блюда" + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetDishesByCategory(ctx context.Context, conn *pgx.Conn, category string, limit, offset int) ([]domain.Dish, error) {
	if category == "" || len(category) == 0 {
		return nil, errors.ValidationError{
			Field:   "category",
			Message: "Категория не может быть пустой",
		}
	}
	limitOffset(limit, offset)
	dish, err := s.dishService.GetByCategoryDish(ctx, conn, category, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetDishesByCategory",
			Message: "Не получилось вернуть блюда по категориям " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetDishesByCuisine(ctx context.Context, conn *pgx.Conn, cuisine string, limit, offset int) ([]domain.Dish, error) {
	if cuisine == "" || len(cuisine) == 0 {
		return nil, errors.ValidationError{
			Field:   "cuisine",
			Message: "Кухня не может быть пустой",
		}
	}
	limitOffset(limit, offset)
	dish, err := s.dishService.GetByCuisineDish(ctx, conn, cuisine, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetDishesByCuisine",
			Message: "Не получилось вернуть блюда по кухням " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetDishesByPriceRange(ctx context.Context, conn *pgx.Conn, max, min, limit, offset int) ([]domain.Dish, error) {
	if max < 0 || min < 0 {
		return nil, errors.ValidationError{
			Field:   "MaxOrMin",
			Message: "Max и min не могут быть меньше нуля",
		}
	}
	limitOffset(limit, offset)
	dish, err := s.dishService.GetByPriceRange(ctx, conn, min, max, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetByPriceRange",
			Message: "Не получилось вернуть блюда в этом диапазоне цен " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetAvailableDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	limitOffset(limit, offset)
	dish, err := s.dishService.GetByAvailable(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetAvailableDishes",
			Message: "Не получилось вернуть блюда доступные блюда " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetVegetarianDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	limitOffset(limit, offset)
	dish, err := s.dishService.GetVegetarianDish(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetVegetarianDishes",
			Message: "Не получилось вернуть вегетарианские блюда " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetVeganDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	limitOffset(limit, offset)
	dish, err := s.dishService.GetVeganDish(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetVeganDishes",
			Message: "Не получилось вернуть веганские блюда " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) GetGlutenFreeDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	limitOffset(limit, offset)
	dish, err := s.dishService.GetGlutenFreeDish(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetGlutenFreeDish",
			Message: "Не получилось вернуть без глютенные блюда " + err.Error(),
		}
	}
	return dish, nil
}

func (s *DishService) UpdateDishService(ctx context.Context, conn *pgx.Conn, id int, dish map[string]interface{}) (*domain.Dish, error) {
	existingDish, err := s.dishService.GetByIDDish(ctx, conn, id)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "dish",
			ID:     id,
		}
	}

	if name, ok := dish["name"].(string); ok {
		existingDish.Name = name
	}
	if description, ok := dish["description"].(string); ok {
		existingDish.Description = description
	}
	if price, ok := dish["price"].(float64); ok {
		existingDish.Price = price
	}
	if category, ok := dish["category"].(string); ok {
		existingDish.Category = category
	}
	if cuisine, ok := dish["cuisine"].(string); ok {
		existingDish.Cuisine = cuisine
	}
	if cookingTime, ok := dish["cooking_time"].(int); ok {
		existingDish.CookingTime = cookingTime
	}
	if isAvailable, ok := dish["is_available"].(bool); ok {
		existingDish.IsAvailable = isAvailable
	}
	if isVegetarian, ok := dish["is_vegetarian"].(bool); ok {
		existingDish.IsVegetarian = isVegetarian
	}
	if isVegan, ok := dish["is_vegan"].(bool); ok {
		existingDish.IsVegan = isVegan
	}
	if isGlutenFree, ok := dish["is_gluten_free"].(bool); ok {
		existingDish.IsGlutenFree = isGlutenFree
	}
	if calories, ok := dish["calories"].(int); ok {
		existingDish.Calories = calories
	}
	if imageURL, ok := dish["image_url"].(string); ok {
		existingDish.ImageURL = imageURL
	}

	if err := existingDish.Validate(); err != nil {
		return nil, errors.ValidationError{
			Field:   "dish",
			Message: err.Error(),
		}
	}

	updated, err := s.dishService.UpdateDish(ctx, conn, id, *existingDish)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrUpdateDish",
			Message: "Не удалось обновить данные у блюда: " + err.Error(),
		}
	}
	return updated, nil
}

func (s *DishService) DeleteDish(ctx context.Context, conn *pgx.Conn, id int) error {
	if id <= 0 {
		return errors.ValidationError{
			Field:   "id",
			Message: "ID не может быть меньше или равен 0",
		}
	}
	if err := s.dishService.DeleteDish(ctx, conn, id); err != nil {
		return errors.BusinessError{
			Code:    "ErrDeleteDish",
			Message: "Не удалось удалить блюдо" + err.Error(),
		}
	}
	return nil
}

func (s *DishService) UpdateAvailability(ctx context.Context, conn *pgx.Conn, id int, availability bool) error {
	if id <= 0 {
		return errors.ValidationError{
			Field:   "id",
			Message: "ID не может быть меньше или равен 0",
		}
	}
	if err := s.dishService.UpdateAvailable(ctx, conn, id, availability); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateAvailability",
			Message: "Не удалось обновить доступность" + err.Error(),
		}
	}
	return nil
}

func (s *DishService) AddDecorator(ctx context.Context, conn *pgx.Conn, id int, decoratorType string) (*domain.Dish, error) {
	dish, err := s.dishService.GetByIDDish(ctx, conn, id)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "dish",
			ID:     id,
		}
	}

	var baseDish decorator.Dish

	switch dish.Category {
	case "pizza":
		baseDish = decorator.NewPizza("средняя")
	case "pasta":
		baseDish = decorator.NewPasta(dish.Name)
	case "salad":
		baseDish = decorator.NewSalad(dish.Name)
	}
	var decoratedDish decorator.Dish

	switch decoratorType {
	case "cheese":
		decoratedDish = decorator.NewCheeseDecorator(baseDish)
	case "bacon":
		decoratedDish = decorator.NewBaconDecorator(baseDish)
	case "extra_sauce":
		decoratedDish = decorator.NewExtraSauceDecorator(baseDish)
	case "gluten_free":
		decoratedDish = decorator.NewGlutenFreeDecorator(baseDish)
	}
	newDescription := decoratedDish.GetDescription()
	newPrice := decoratedDish.GetPrice()

	updatedDish := &domain.Dish{
		DishID:       dish.DishID,
		Name:         dish.Name,
		Description:  newDescription,
		Price:        newPrice,
		Category:     dish.Category,
		Cuisine:      dish.Cuisine,
		CookingTime:  dish.CookingTime,
		IsAvailable:  dish.IsAvailable,
		IsVegetarian: dish.IsVegetarian,
		IsVegan:      dish.IsVegan,
		IsGlutenFree: dish.IsGlutenFree,
		Calories:     dish.Calories,
		ImageURL:     dish.ImageURL,
	}

	result, err := s.dishService.UpdateDish(ctx, conn, id, *updatedDish)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrUpdateDish",
			Message: "Не удалось обновить блюдо: " + err.Error(),
		}
	}
	return result, nil
}

func (s *DishService) RemoveDecorator(ctx context.Context, conn *pgx.Conn, dishID int, decoratorType string) (*domain.Dish, error) {
	dish, err := s.dishService.GetByIDDish(ctx, conn, dishID)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "dish",
			ID:     dishID,
		}
	}
	newDescription := s.removeDecoratorFromDescription(dish.Description, decoratorType)

	basePrice := s.getBasePrice(dish, decoratorType)

	updatedDish := &domain.Dish{
		DishID:       dish.DishID,
		Name:         dish.Name,
		Description:  newDescription,
		Price:        basePrice,
		Category:     dish.Category,
		Cuisine:      dish.Cuisine,
		CookingTime:  dish.CookingTime,
		IsAvailable:  dish.IsAvailable,
		IsVegetarian: dish.IsVegetarian,
		IsVegan:      dish.IsVegan,
		IsGlutenFree: dish.IsGlutenFree,
		Calories:     dish.Calories,
		ImageURL:     dish.ImageURL,
	}

	result, err := s.dishService.UpdateDish(ctx, conn, dishID, *updatedDish)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrRemoveDecorator",
			Message: "Не удалось обновить блюдо " + err.Error(),
		}
	}
	return result, nil
}

func (s *DishService) removeDecoratorFromDescription(description, decoratorType string) string {
	decoratorMap := map[string]string{
		"cheese":      "+ сыр",
		"bacon":       "+ бекон",
		"extra_sauce": "+ соус",
		"gluten_free": "(без глютена)",
	}

	decoratorText, ok := decoratorMap[decoratorType]
	if !ok {
		return description
	}

	result := ""
	for i := 0; i < len(description); {
		if i+len(decoratorText) <= len(description) && description[i:i+len(decoratorText)] == decoratorText {
			i += len(decoratorText)
			for i < len(description) && description[i] == ' ' {
				i++
			}
		} else {
			result += string(description[i])
			i++
		}
	}

	return result
}

func (s *DishService) getBasePrice(dish *domain.Dish, decoratorType string) float64 {
	decoratorPrices := map[string]float64{
		"cheese":      50.0,
		"bacon":       80.0,
		"extra_sauce": 30.0,
		"gluten_free": 100.0,
	}

	price, ok := decoratorPrices[decoratorType]
	if !ok {
		return dish.Price
	}

	basePrice := dish.Price - price
	if basePrice < 0 {
		basePrice = dish.Price
	}
	return basePrice
}

func (s *DishService) GetMenuTree(ctx context.Context, conn *pgx.Conn) (*composite.MenuCategory, error) {
	dishes, err := s.dishService.GetAllDishes(ctx, conn, 1000, 0)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetDishes",
			Message: "Не удалось получить блюда для меню: " + err.Error(),
		}
	}

	root := composite.NewMenuCategory("Меню ресторана", "Все блюда ресторана")

	categories := make(map[string]*composite.MenuCategory)
	for _, dish := range dishes {
		categoryName := dish.Category
		if _, exists := categories[categoryName]; !exists {
			categories[categoryName] = composite.NewMenuCategory(categoryName, "Блюда категории "+categoryName)
		}
	}
	for _, category := range categories {
		root.Add(category)
	}

	for _, dish := range dishes {
		category := categories[dish.Category]
		if category == nil {
			dishComponent := composite.NewDishComponent(
				dish.Name,
				dish.Description,
				dish.Category,
				dish.Cuisine,
				dish.Price,
				dish.CookingTime,
				[]string{},
			)
			category.Add(dishComponent)
		}
	}
	return root, nil
}

func (s *DishService) GetComboMeal(ctx context.Context, conn *pgx.Conn, comboType string) (*composite.ComboMeal, error) {
	var factory abstract_factory.MenuFactory

	switch comboType {
	case "business":
		factory = abstract_factory.NewBusinessLunchFactory()
	case "kids":
		factory = abstract_factory.NewKidsMenuFactory()
	case "vegan":
		factory = abstract_factory.NewVeganMenuFactory()
	default:
		return nil, errors.ValidationError{
			Field:   "combo_type",
			Message: fmt.Sprintf("Неизвестный тип комбо: %s", comboType),
		}
	}

	menuSet := factory.CreateMenuSet()

	comboMeal := composite.NewComboMeal(
		menuSet.Name,
		menuSet.Description,
		menuSet.Discount,
	)

	dishNames := menuSet.GetItems()

	for _, dishName := range dishNames {
		dishes, err := s.dishService.GetByNameDish(ctx, conn, dishName, 1, 0)
		if err != nil || len(dishes) == 0 {
			continue
		}

		dish := dishes[0]

		dishComponent := composite.NewDishComponent(
			dish.Name,
			dish.Description,
			dish.Category,
			dish.Cuisine,
			dish.Price,
			dish.CookingTime,
			[]string{},
		)
		comboMeal.Add(dishComponent)
	}

	return comboMeal, nil
}

func (s *DishService) SearchDishes(ctx context.Context, conn *pgx.Conn, query string, limit, offset int) ([]domain.Dish, int, error) {
	limit, offset = limitOffset(limit, offset)

	if query == "" {
		return nil, 0, errors.ValidationError{
			Field:   "query",
			Message: "Поисковый запрос не может быть пустым",
		}
	}

	var allDishes []domain.Dish
	var err error

	if s.menuProxy != nil {
		menuItems := s.menuProxy.GetAllItems()
		allDishes = make([]domain.Dish, len(menuItems))
		for i, item := range menuItems {
			allDishes[i] = domain.Dish{
				DishID:   item.ID,
				Name:     item.Name,
				Price:    item.Price,
				Category: item.Category,
			}
		}
	} else {
		allDishes, err = s.dishService.GetAllDishes(ctx, conn, 1000, 0)
		if err != nil {
			return nil, 0, errors.BusinessError{
				Code:    "ErrGetDishes",
				Message: "Не удалось получить блюда для поиска: " + err.Error(),
			}
		}
	}

	var results []domain.Dish
	queryLower := toLower(query)

	for _, dish := range allDishes {
		if containsIgnoreCase(dish.Name, queryLower) {
			results = append(results, dish)
		}
	}

	total := len(results)
	start := offset
	if start > total {
		start = total
	}
	end := offset + limit
	if end > total {
		end = total
	}

	paginatedResults := results[start:end]

	return paginatedResults, total, nil
}

func toLower(str string) string {
	result := ""
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			result += string(ch + 32)
		} else {
			result += string(ch)
		}
	}
	return result
}

func containsIgnoreCase(str, substr string) bool {
	strLower := toLower(str)
	substrLower := toLower(substr)
	if len(substrLower) == 0 {
		return true
	}
	for i := 0; i <= len(strLower)-len(substrLower); i++ {
		if strLower[i:i+len(substrLower)] == substrLower {
			return true
		}
	}
	return false
}
