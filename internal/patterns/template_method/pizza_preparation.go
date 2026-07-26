package template_method

import "fmt"

// приготовление пиццы
type PizzaPreparation struct {
	*DishPreparation
	Size    string
	Topping string
}

func NewPizzaPreparation(size, topping string) *PizzaPreparation {
	return &PizzaPreparation{
		DishPreparation: NewDishPreparation("Пицца " + topping),
		Size:            size,
		Topping:         topping,
	}
}

// сбор ингредиентов для пиццы
func (p *PizzaPreparation) GatherIngredients() {
	fmt.Printf("Шаг 1: Сбор ингредиентов для пиццы %s\n", p.Size)
	fmt.Println("Мука, вода, дрожжи (тесто)")
	fmt.Printf("Томатный соус, сыр Моцарелла, %s\n", p.Topping)
	fmt.Println("Оливковое масло, базилик, орегано")
}

// PrepareIngredients — подготовка ингредиентов для пиццы
func (p *PizzaPreparation) PrepareIngredients() {
	fmt.Println("Шаг 2: Подготовка ингредиентов")
	fmt.Println("Замес теста для пиццы")
	fmt.Printf("Раскатка теста до размера %s\n", p.Size)
	fmt.Println("Нарезка ингредиентов")
}

// приготовление пиццы
func (p *PizzaPreparation) Cook() {
	fmt.Printf("Шаг 3: Выпекание пиццы в печи\n")
	fmt.Printf("Температура: 250°C\n")
	fmt.Printf("Время: 15 минут\n")
	fmt.Printf("Проверка готовности: золотистая корочка\n")
}

// подача пиццы на тарелку
func (p *PizzaPreparation) Plate() {
	fmt.Println("Шаг 4: Подача на тарелку")
	fmt.Println("Круглая деревянная доска")
	fmt.Println("Нарезана на 8 кусков")
}

// украшение пиццы
func (p *PizzaPreparation) Garnish() {
	fmt.Println("Шаг 5: Украшение пиццы")
	fmt.Println("Листики свежего базилика")
	fmt.Println("Оливковое масло по краям")
}

// время приготовления пиццы
func (p *PizzaPreparation) GetCookingTime() int {
	return 20
}
