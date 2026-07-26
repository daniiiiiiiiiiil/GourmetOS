package decorator

type Dish interface {
	GetDescription() string
	GetPrice() float64
}

type DishDecorator struct {
	Dish Dish
}

func (d *DishDecorator) GetDescription() string {
	return d.Dish.GetDescription()
}

func (d *DishDecorator) GetPrice() float64 {
	return d.Dish.GetPrice()
}
