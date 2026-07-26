package decorator

type ExtraSauceDecorator struct {
	DishDecorator
}

func NewExtraSauceDecorator(dish Dish) *ExtraSauceDecorator {
	return &ExtraSauceDecorator{
		DishDecorator: DishDecorator{Dish: dish},
	}
}

// добавляет " + соус" к описанию
func (e *ExtraSauceDecorator) GetDescription() string {
	return e.Dish.GetDescription() + " + соус"
}

// добавляет стоимость соуса
func (e *ExtraSauceDecorator) GetPrice() float64 {
	return e.Dish.GetPrice() + 30.0
}
