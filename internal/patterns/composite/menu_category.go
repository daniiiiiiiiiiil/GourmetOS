package composite

import "fmt"

// категория меню (композит)
// Может содержать другие категории или блюда
type MenuCategory struct {
	BaseComponent
	children []MenuComponent
}

func NewMenuCategory(name, description string) *MenuCategory {
	return &MenuCategory{
		BaseComponent: BaseComponent{
			name:        name,
			description: description,
			price:       0,
		},
		children: []MenuComponent{},
	}
}

// добавляет дочерний компонент
func (m *MenuCategory) Add(child MenuComponent) {
	m.children = append(m.children, child)
}

// удаляет дочерний компонент
func (m *MenuCategory) Remove(child MenuComponent) {
	for i, c := range m.children {
		if c == child {
			m.children = append(m.children[:i], m.children[i+1:]...)
			break
		}
	}
}

// возвращает дочерний компонент по индексу
func (m *MenuCategory) GetChild(index int) MenuComponent {
	if index < 0 || index >= len(m.children) {
		return nil
	}
	return m.children[index]
}

// возвращает сумму цен всех дочерних элементов
func (m *MenuCategory) GetPrice() float64 {
	total := 0.0
	for _, child := range m.children {
		total += child.GetPrice()
	}
	return total
}

// возвращает количество дочерних элементов
func (m *MenuCategory) GetChildrenCount() int {
	return len(m.children)
}

// возвращает все блюда в категории (рекурсивно)
func (m *MenuCategory) GetAllDishes() []*DishComponent {
	result := []*DishComponent{}
	for _, child := range m.children {
		if dish, ok := child.(*DishComponent); ok {
			result = append(result, dish)
		}
		if category, ok := child.(*MenuCategory); ok {
			result = append(result, category.GetAllDishes()...)
		}
	}
	return result
}

// выводит информацию о категории и всех дочерних элементах
func (m *MenuCategory) Print(indent string) {
	fmt.Printf("%s %s\n", indent, m.name)
	if m.description != "" {
		fmt.Printf("%s   Описание: %s\n", indent, m.description)
	}
	fmt.Printf("%s   Количество: %d позиций\n", indent, len(m.children))

	for _, child := range m.children {
		child.Print(indent + "  ")
	}
}

// IsComposite — всегда true (это композит)
func (m *MenuCategory) IsComposite() bool {
	return true
}
