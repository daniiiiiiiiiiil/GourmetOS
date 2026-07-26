package abstract_factory

type BusinessLunchFactory struct{}

func NewBusinessLunchFactory() *BusinessLunchFactory {
	return &BusinessLunchFactory{}
}

// создает основное блюдо для бизнес-ланча
func (b *BusinessLunchFactory) CreateMainDish() *MainDish {
	return &MainDish{
		Name:        "Стейк с овощами",
		Description: "Говяжий стейк с запеченными овощами",
		Price:       450.0,
		Weight:      250,
		CookingTime: 25,
	}
}

// создает гарнир для бизнес-ланча
func (b *BusinessLunchFactory) CreateSideDish() *SideDish {
	return &SideDish{
		Name:        "Картофель фри",
		Description: "Хрустящий картофель со специями",
		Price:       120.0,
		Weight:      150,
	}
}

// создает напиток для бизнес-ланча
func (b *BusinessLunchFactory) CreateDrink() *Drink {
	return &Drink{
		Name:        "Кофе американо",
		Description: "Свежесваренный кофе",
		Price:       80.0,
		Volume:      200,
	}
}

// создает десерт для бизнес-ланча
func (b *BusinessLunchFactory) CreateDessert() *Dessert {
	return &Dessert{
		Name:        "Тирамису",
		Description: "Итальянский десерт с маскарпоне",
		Price:       180.0,
		Weight:      120,
		Calories:    350,
	}
}

// создает комплексный обед
func (b *BusinessLunchFactory) CreateMenuSet() *MenuSet {
	return NewMenuSet(
		"Бизнес-ланч",
		"Сытный обед для деловых людей",
		b.CreateMainDish(),
		b.CreateSideDish(),
		b.CreateDrink(),
		b.CreateDessert(),
		15.0,
	)
}

func (b *BusinessLunchFactory) GetName() string {
	return "Бизнес-ланч"
}

func (b *BusinessLunchFactory) GetDescription() string {
	return "Комплексный обед для деловых людей со скидкой 15%"
}
