package decorator

type GlutenFreeDecorator struct {
	DishDecorator
}

func NewGlutenFreeDecorator(dish Dish) *GlutenFreeDecorator {
	return &GlutenFreeDecorator{
		DishDecorator: DishDecorator{Dish: dish},
	}
}

// добавляет " (без глютена)" к описанию
func (g *GlutenFreeDecorator) GetDescription() string {
	return g.Dish.GetDescription() + " (без глютена)"
}

// добавляет наценку за безглютеновое тесто
func (g *GlutenFreeDecorator) GetPrice() float64 {
	return g.Dish.GetPrice() + 100.0
}
