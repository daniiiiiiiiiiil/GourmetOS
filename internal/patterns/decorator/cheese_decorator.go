package decorator

type CheeseDecorator struct {
	DishDecorator
}

func NewCheeseDecorator(dish Dish) *CheeseDecorator {
	return &CheeseDecorator{
		DishDecorator: DishDecorator{Dish: dish},
	}
}

// добавляет " + сыр" к описанию
func (c *CheeseDecorator) GetDescription() string {
	return c.Dish.GetDescription() + " + сыр"
}

// добавляет стоимость сыра
func (c *CheeseDecorator) GetPrice() float64 {
	return c.Dish.GetPrice() + 50.0
}
