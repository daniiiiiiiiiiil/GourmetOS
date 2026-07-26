package abstract_factory

type VeganMenuFactory struct{}

func NewVeganMenuFactory() *VeganMenuFactory {
	return &VeganMenuFactory{}
}

// создает основное блюдо для веганского меню
func (v *VeganMenuFactory) CreateMainDish() *MainDish {
	return &MainDish{
		Name:        "Бургер с нутовой котлетой",
		Description: "Веганский бургер с овощами и соусом",
		Price:       380.0,
		Weight:      280,
		CookingTime: 20,
	}
}

// создает гарнир для веганского меню
func (v *VeganMenuFactory) CreateSideDish() *SideDish {
	return &SideDish{
		Name:        "Овощной салат",
		Description: "Свежие овощи с оливковым маслом",
		Price:       130.0,
		Weight:      180,
	}
}

// создает напиток для веганского меню
func (v *VeganMenuFactory) CreateDrink() *Drink {
	return &Drink{
		Name:        "Смузи",
		Description: "Банан, шпинат, яблоко",
		Price:       110.0,
		Volume:      300,
	}
}

// создает десерт для веганского меню
func (v *VeganMenuFactory) CreateDessert() *Dessert {
	return &Dessert{
		Name:        "Ореховый торт",
		Description: "Без сахара, без глютена, без молока",
		Price:       200.0,
		Weight:      130,
		Calories:    280,
	}
}

// создает веганский набор
func (v *VeganMenuFactory) CreateMenuSet() *MenuSet {
	return NewMenuSet(
		"Веганское меню",
		"100% растительные блюда",
		v.CreateMainDish(),
		v.CreateSideDish(),
		v.CreateDrink(),
		v.CreateDessert(),
		12.0,
	)
}

func (v *VeganMenuFactory) GetName() string {
	return "Веганское меню"
}

func (v *VeganMenuFactory) GetDescription() string {
	return "Полноценное питание без продуктов животного происхождения со скидкой 12%"
}
