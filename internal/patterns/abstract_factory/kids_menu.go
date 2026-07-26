package abstract_factory

type KidsMenuFactory struct{}

// конструктор
func NewKidsMenuFactory() *KidsMenuFactory {
	return &KidsMenuFactory{}
}

// создает основное блюдо для детского меню
func (k *KidsMenuFactory) CreateMainDish() *MainDish {
	return &MainDish{
		Name:        "Куриные наггетсы",
		Description: "Нежные куриные кусочки в панировке",
		Price:       280.0,
		Weight:      180,
		CookingTime: 15,
	}
}

// создает гарнир для детского меню
func (k *KidsMenuFactory) CreateSideDish() *SideDish {
	return &SideDish{
		Name:        "Картошка фри",
		Description: "Маленькая порция картошки фри",
		Price:       70.0,
		Weight:      100,
	}
}

// создает напиток для детского меню
func (k *KidsMenuFactory) CreateDrink() *Drink {
	return &Drink{
		Name:        "Сок яблочный",
		Description: "Натуральный яблочный сок",
		Price:       60.0,
		Volume:      200,
	}
}

// создает десерт для детского меню
func (k *KidsMenuFactory) CreateDessert() *Dessert {
	return &Dessert{
		Name:        "Мороженое",
		Description: "Ванильное мороженое с шоколадной крошкой",
		Price:       90.0,
		Weight:      100,
		Calories:    200,
	}
}

// создает детский набор
func (k *KidsMenuFactory) CreateMenuSet() *MenuSet {
	return NewMenuSet(
		"Детское меню",
		"Вкусный и полезный набор для детей",
		k.CreateMainDish(),
		k.CreateSideDish(),
		k.CreateDrink(),
		k.CreateDessert(),
		10.0,
	)
}

func (k *KidsMenuFactory) GetName() string {
	return "Детское меню"
}

func (k *KidsMenuFactory) GetDescription() string {
	return "Специальное меню для детей со скидкой 10%"
}
