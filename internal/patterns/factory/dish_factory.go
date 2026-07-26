package factory

//главная суть new будет в сервисе

type Dish interface {
	GetName() string
	GetPrice() float64
	GetIngredients() []string
	GetCookingTime() int
}

type DishFactory interface {
	CreatePizza() Dish
	CreatePasta() Dish
	CreateSalad() Dish
	CreateDrink() Dish
	GetKitchenName() string
}

type BaseDish struct {
	Name        string
	Price       float64
	Ingredients []string
	CookingTime int
}

func (b BaseDish) GetName() string {
	return b.Name
}

func (b BaseDish) GetPrice() float64 {
	return b.Price
}

func (b BaseDish) GetIngredients() []string {
	return b.Ingredients
}

func (b BaseDish) GetCookingTime() int {
	return b.CookingTime
}
