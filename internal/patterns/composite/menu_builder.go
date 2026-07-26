package composite

type MenuBuilder struct {
	root *MenuCategory
}

func NewMenuBuilder() *MenuBuilder {
	return &MenuBuilder{}
}

// строит полное меню ресторана
func (b *MenuBuilder) BuildFullMenu() *MenuCategory {
	// Корневая категория
	root := NewMenuCategory("Меню ресторана", "Все блюда ресторана")

	// Итальянская кухня
	italian := NewMenuCategory("Итальянская кухня", "Блюда итальянской кухни")

	pizza := NewMenuCategory("Пицца", "Итальянская пицца")
	pizza.Add(NewDishComponent(
		"Маргарита",
		"Томатный соус, моцарелла, базилик",
		"pizza", "italian",
		450.0, 15,
		[]string{"мука", "томатный соус", "моцарелла", "базилик"},
	))
	pizza.Add(NewDishComponent(
		"Пепперони",
		"Томатный соус, моцарелла, пепперони",
		"pizza", "italian",
		520.0, 15,
		[]string{"мука", "томатный соус", "моцарелла", "пепперони"},
	))

	pasta := NewMenuCategory("Паста", "Итальянская паста")
	pasta.Add(NewDishComponent(
		"Карбонара",
		"Спагетти, бекон, яйцо, пармезан",
		"pasta", "italian",
		380.0, 20,
		[]string{"спагетти", "бекон", "яйцо", "пармезан"},
	))
	pasta.Add(NewDishComponent(
		"Болоньезе",
		"Спагетти, мясной соус, пармезан",
		"pasta", "italian",
		420.0, 25,
		[]string{"спагетти", "говяжий фарш", "томаты", "пармезан"},
	))

	italian.Add(pizza)
	italian.Add(pasta)

	// Японская кухня
	japanese := NewMenuCategory("Японская кухня", "Блюда японской кухни")
	japanese.Add(NewDishComponent(
		"Окономияки",
		"Японская пицца с капустой и свининой",
		"pizza", "japanese",
		500.0, 25,
		[]string{"мука", "капуста", "яйцо", "свинина", "соус"},
	))
	japanese.Add(NewDishComponent(
		"Рамен",
		"Лапша в свином бульоне с яйцом и нори",
		"pasta", "japanese",
		420.0, 30,
		[]string{"лапша", "свиной бульон", "яйцо", "нори"},
	))

	combo := NewComboMeal(
		"Бизнес-ланч",
		"Комплексный обед с 20% скидкой",
		20.0,
	)
	combo.Add(NewDishComponent(
		"Маргарита",
		"Томатный соус, моцарелла, базилик",
		"pizza", "italian",
		450.0, 15,
		[]string{"мука", "томатный соус", "моцарелла", "базилик"},
	))
	combo.Add(NewDishComponent(
		"Салат Капрезе",
		"Томаты, моцарелла, базилик",
		"salad", "italian",
		280.0, 10,
		[]string{"томаты", "моцарелла", "базилик"},
	))
	combo.Add(NewDishComponent(
		"Лимончелло",
		"Освежающий лимонный напиток",
		"drink", "italian",
		150.0, 5,
		[]string{"лимон", "сахар", "вода"},
	))

	// Собираем всё вместе
	root.Add(italian)
	root.Add(japanese)
	root.Add(combo)

	b.root = root
	return root
}

// возвращает корневой элемент
func (b *MenuBuilder) GetRoot() *MenuCategory {
	return b.root
}

// ищет блюдо по названию
func (b *MenuBuilder) FindDishByName(name string) *DishComponent {
	if b.root == nil {
		return nil
	}
	return b.findDishRecursive(b.root, name)
}

// рекурсивный поиск блюда
func (b *MenuBuilder) findDishRecursive(component MenuComponent, name string) *DishComponent {
	// Если это блюдо и имя совпадает
	if dish, ok := component.(*DishComponent); ok {
		if dish.GetName() == name {
			return dish
		}
	}

	// Если это композит, ищем внутри
	if component.IsComposite() {
		// Перебираем всех детей
		if category, ok := component.(*MenuCategory); ok {
			for i := 0; i < category.GetChildrenCount(); i++ {
				child := category.GetChild(i)
				if result := b.findDishRecursive(child, name); result != nil {
					return result
				}
			}
		}
		if combo, ok := component.(*ComboMeal); ok {
			for i := 0; i < len(combo.children); i++ {
				child := combo.GetChild(i)
				if result := b.findDishRecursive(child, name); result != nil {
					return result
				}
			}
		}
	}
	return nil
}
