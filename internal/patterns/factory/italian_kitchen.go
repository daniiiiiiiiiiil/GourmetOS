package factory

type ItalianKitchen struct{}

func NewItalianKitchen() *ItalianKitchen {
	return &ItalianKitchen{}
}

func (i *ItalianKitchen) CreatePizza() Dish {
	return &BaseDish{
		Name:        "Маргарита",
		Price:       450.0,
		Ingredients: []string{"мука", "томатный соус", "моцарелла", "базилик"},
		CookingTime: 15,
	}
}

func (i *ItalianKitchen) CreatePasta() Dish {
	return &BaseDish{
		Name:        "Карбонара",
		Price:       380.0,
		Ingredients: []string{"спагетти", "бекон", "яйцо", "пармезан"},
		CookingTime: 20,
	}
}

func (i *ItalianKitchen) CreateSalad() Dish {
	return &BaseDish{
		Name:        "Капрезе",
		Price:       280.0,
		Ingredients: []string{"томаты", "моцарелла", "базилик", "оливковое масло"},
		CookingTime: 10,
	}
}

func (i *ItalianKitchen) CreateDrink() Dish {
	return &BaseDish{
		Name:        "Лимончелло",
		Price:       150.0,
		Ingredients: []string{"лимон", "сахар", "вода", "лед"},
		CookingTime: 5,
	}
}

func (i *ItalianKitchen) GetKitchenName() string {
	return "Итальянская кухня"
}
