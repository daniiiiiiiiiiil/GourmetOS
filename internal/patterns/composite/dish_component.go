package composite

import "fmt"

// блюдо (лист в дереве меню)
// Не может содержать дочерние элементы
type DishComponent struct {
	BaseComponent
	category    string
	cuisine     string
	cookingTime int
	ingredients []string
}

func NewDishComponent(name, description, category, cuisine string, price float64, cookingTime int, ingredients []string) *DishComponent {
	return &DishComponent{
		BaseComponent: BaseComponent{
			name:        name,
			price:       price,
			description: description,
		},
		category:    category,
		cuisine:     cuisine,
		cookingTime: cookingTime,
		ingredients: ingredients,
	}
}

func (d *DishComponent) GetCategory() string {
	return d.category
}

func (d *DishComponent) GetCuisine() string {
	return d.cuisine
}

func (d *DishComponent) GetCookingTime() int {
	return d.cookingTime
}

func (d *DishComponent) GetIngredients() []string {
	return d.ingredients
}

// выводит информацию о блюде
func (d *DishComponent) Print(indent string) {
	fmt.Printf("%s %s\n", indent, d.name)
	fmt.Printf("%s   Цена: %.2f руб.\n", indent, d.price)
	fmt.Printf("%s   Описание: %s\n", indent, d.description)
	fmt.Printf("%s   Категория: %s\n", indent, d.category)
	fmt.Printf("%s   Кухня: %s\n", indent, d.cuisine)
	fmt.Printf("%s   Время: %d мин.\n", indent, d.cookingTime)
	fmt.Printf("%s   Ингредиенты: %v\n", indent, d.ingredients)
}

// всегда false (это лист)
func (d *DishComponent) IsComposite() bool {
	return false
}
