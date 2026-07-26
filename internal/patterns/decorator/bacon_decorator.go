package decorator

type BaconDecorator struct {
	DishDecorator
}

func NewBaconDecorator(dish Dish) *BaconDecorator {
	return &BaconDecorator{
		DishDecorator: DishDecorator{Dish: dish},
	}
}

// добавляет " + бекон" к описанию
func (b *BaconDecorator) GetDescription() string {
	return b.Dish.GetDescription() + " + бекон"
}

// добавляет стоимость бекона
func (b *BaconDecorator) GetPrice() float64 {
	return b.Dish.GetPrice() + 80.0
}
