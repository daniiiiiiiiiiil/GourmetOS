package factory

type MexicanKitchen struct{}

func NewMexicanKitchen() *MexicanKitchen {
	return &MexicanKitchen{}
}

func (m *MexicanKitchen) CreatePizza() Dish {
	return &BaseDish{
		Name:        "Тортилья",
		Price:       470.0,
		Ingredients: []string{"тортилья", "фасоль", "сыр", "сальса", "гуакамоле"},
		CookingTime: 20,
	}
}

func (m *MexicanKitchen) CreatePasta() Dish {
	return &BaseDish{
		Name:        "Чили кон карне",
		Price:       450.0,
		Ingredients: []string{"говядина", "фасоль", "томаты", "перец чили", "лук"},
		CookingTime: 35,
	}
}

func (m *MexicanKitchen) CreateSalad() Dish {
	return &BaseDish{
		Name:        "Салат с кактусом",
		Price:       300.0,
		Ingredients: []string{"кактус", "помидор", "лук", "авокадо", "сыр"},
		CookingTime: 15,
	}
}

func (m *MexicanKitchen) CreateDrink() Dish {
	return &BaseDish{
		Name:        "Агуа Фреска",
		Price:       130.0,
		Ingredients: []string{"вода", "фрукты", "сахар", "лед"},
		CookingTime: 5,
	}
}

func (m *MexicanKitchen) GetKitchenName() string {
	return "Мексиканская кухня"
}
