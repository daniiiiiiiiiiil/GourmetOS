package composite

import "fmt"

// комбо-меню (композит)
// Объединяет несколько блюд в один набор со скидкой
type ComboMeal struct {
	BaseComponent
	children        []MenuComponent
	discountPercent float64
}

func NewComboMeal(name, description string, discountPercent float64) *ComboMeal {
	return &ComboMeal{
		BaseComponent: BaseComponent{
			name:        name,
			description: description,
			price:       0,
		},
		children:        []MenuComponent{},
		discountPercent: discountPercent,
	}
}

func (c *ComboMeal) Add(child MenuComponent) {
	c.children = append(c.children, child)
}

func (c *ComboMeal) Remove(child MenuComponent) {
	for i, comp := range c.children {
		if comp == child {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}
	}
}

func (c *ComboMeal) GetChild(index int) MenuComponent {
	if index < 0 || index >= len(c.children) {
		return nil
	}
	return c.children[index]
}

// возвращает сумму цен всех блюд со скидкой
func (c *ComboMeal) GetPrice() float64 {
	total := 0.0
	for _, child := range c.children {
		total += child.GetPrice()
	}
	// Применяем скидку
	discount := total * (c.discountPercent / 100)
	return total - discount
}

// возвращает сумму без скидки
func (c *ComboMeal) GetOriginalPrice() float64 {
	total := 0.0
	for _, child := range c.children {
		total += child.GetPrice()
	}
	return total
}

func (c *ComboMeal) GetDiscountPercent() float64 {
	return c.discountPercent
}

// возвращает список блюд в комбо
func (c *ComboMeal) GetItems() []string {
	result := []string{}
	for _, child := range c.children {
		result = append(result, child.GetName())
	}
	return result
}

func (c *ComboMeal) Print(indent string) {
	fmt.Printf("%s %s\n", indent, c.name)
	fmt.Printf("%s   Описание: %s\n", indent, c.description)
	fmt.Printf("%s   Скидка: %.0f%%\n", indent, c.discountPercent)
	fmt.Printf("%s   Сумма без скидки: %.2f руб.\n", indent, c.GetOriginalPrice())
	fmt.Printf("%s   Итоговая цена: %.2f руб.\n", indent, c.GetPrice())
	fmt.Printf("%s   Экономия: %.2f руб.\n", indent, c.GetOriginalPrice()-c.GetPrice())
	fmt.Printf("%s   Включает:\n", indent)

	for _, child := range c.children {
		child.Print(indent + "     ")
	}
}

// всегда true (это композит)
func (c *ComboMeal) IsComposite() bool {
	return true
}
