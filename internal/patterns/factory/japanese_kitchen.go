package factory

type JapaneseKitchen struct{}

func NewJapaneseKitchen() *JapaneseKitchen {
	return &JapaneseKitchen{}
}

func (j *JapaneseKitchen) CreatePizza() Dish {
	return &BaseDish{
		Name:        "Окономияки",
		Price:       500.0,
		Ingredients: []string{"мука", "капуста", "яйцо", "свинина", "соус"},
		CookingTime: 25,
	}
}

func (j *JapaneseKitchen) CreatePasta() Dish {
	return &BaseDish{
		Name:        "Рамен",
		Price:       420.0,
		Ingredients: []string{"лапша", "свиной бульон", "яйцо", "нори", "свинина"},
		CookingTime: 30,
	}
}

func (j *JapaneseKitchen) CreateSalad() Dish {
	return &BaseDish{
		Name:        "Вакамэ",
		Price:       250.0,
		Ingredients: []string{"водоросли", "огурец", "кунжут", "соевый соус"},
		CookingTime: 10,
	}
}

func (j *JapaneseKitchen) CreateDrink() Dish {
	return &BaseDish{
		Name:        "Маття",
		Price:       180.0,
		Ingredients: []string{"зеленый чай", "молоко", "сахар"},
		CookingTime: 5,
	}
}

func (j *JapaneseKitchen) GetKitchenName() string {
	return "Японская кухня"
}
