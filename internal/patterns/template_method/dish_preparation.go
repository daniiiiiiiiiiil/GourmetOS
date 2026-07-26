package template_method

import "fmt"

// базовый класс с шаблонным методом
// Определяет алгоритм приготовления блюда
// Подклассы переопределяют конкретные шаги
type DishPreparation struct {
	DishName string
}

func NewDishPreparation(name string) *DishPreparation {
	return &DishPreparation{DishName: name}
}

//	шаблонный метод (алгоритм приготовления)
//
// Это "скелет" алгоритма, который нельзя изменить
func (d *DishPreparation) Prepare() {
	fmt.Printf("\nПриготовление: %s\n", d.DishName)
	fmt.Println("=====================================")

	// 1. Подготовка ингредиентов (шаг 1)
	d.GatherIngredients()

	// 2. Нарезка и подготовка (шаг 2)
	d.PrepareIngredients()

	// 3. Основное приготовление (шаг 3)
	d.Cook()

	// 4. Подача на тарелку (шаг 4)
	d.Plate()

	// 5. Украшение (шаг 5)
	d.Garnish()

	d.Serve()

	fmt.Println("=====================================")
	fmt.Printf("%s готово\n", d.DishName)
}

// собирает ингредиенты (общий шаг)
// Может быть переопределен в подклассах
func (d *DishPreparation) GatherIngredients() {
	fmt.Println("Шаг 1: Сбор ингредиентов")
}

// подготавливает ингредиенты (общий шаг)
// Может быть переопределен в подклассах
func (d *DishPreparation) PrepareIngredients() {
	fmt.Println("Шаг 2: Нарезка и подготовка")
}

// основное приготовление (ОБЯЗАТЕЛЬНО переопределяется)
// Подклассы должны реализовать свой способ приготовления
func (d *DishPreparation) Cook() {
	fmt.Println("Шаг 3: Приготовление (базовый метод)")
}

// подача на тарелку (ОБЯЗАТЕЛЬНО переопределяется)
func (d *DishPreparation) Plate() {
	fmt.Println("Шаг 4: Подача на тарелку (базовый метод)")
}

// украшение (может быть переопределен или оставлен по умолчанию)
func (d *DishPreparation) Garnish() {
	fmt.Println("Украшение зеленью (стандартно)")
}

// подача клиенту (общий финальный шаг)
// Обычно не переопределяется
func (d *DishPreparation) Serve() {
	fmt.Println("   🤵 Шаг 6: Подача клиенту")
}

// возвращает время приготовления (базовый метод)
// Подклассы переопределяют для указания своего времени
func (d *DishPreparation) GetCookingTime() int {
	return 15
}
