package template_method

import "fmt"

// приготовление салата
type SaladPreparation struct {
	*DishPreparation
	Type string // "цезарь", "греческий", "капрезе"
}

func NewSaladPreparation(saladType string) *SaladPreparation {
	return &SaladPreparation{
		DishPreparation: NewDishPreparation("Салат " + saladType),
		Type:            saladType,
	}
}

// сбор ингредиентов для салата
func (s *SaladPreparation) GatherIngredients() {
	fmt.Printf("Шаг 1: Сбор ингредиентов для салата %s\n", s.Type)
	if s.Type == "цезарь" {
		fmt.Println("Ромейн (салат), курица, пармезан")
		fmt.Println("Гренки, соус Цезарь")
	} else if s.Type == "греческий" {
		fmt.Println("Помидоры, огурцы, перец, лук")
		fmt.Println("Маслины, сыр Фета")
	} else {
		fmt.Println("Помидоры, Моцарелла, базилик")
		fmt.Println("Оливковое масло, бальзамик")
	}
}

// подготовка ингредиентов для салата
func (s *SaladPreparation) PrepareIngredients() {
	fmt.Println("Шаг 2: Подготовка ингредиентов")
	fmt.Println("Мытье и сушка овощей")
	if s.Type == "цезарь" {
		fmt.Println("Жарка курицы на гриле")
		fmt.Println("Приготовление гренок")
	} else if s.Type == "греческий" {
		fmt.Println("Нарезка овощей крупными кусками")
	} else {
		fmt.Println("Нарезка помидоров и моцареллы")
	}
}

// приготовление салата (особенность: тепловая обработка не требуется)
func (s *SaladPreparation) Cook() {
	fmt.Println("Шаг 3: Сборка салата")
	if s.Type == "цезарь" {
		fmt.Println("Смешивание салата с соусом")
		fmt.Println("Добавление гренок и курицы")
	} else if s.Type == "греческий" {
		fmt.Println("Смешивание овощей с маслом")
		fmt.Println("Добавление маслин и Феты")
	} else {
		fmt.Println("Выкладывание кружочками")
		fmt.Println("Поливка оливковым маслом")
		fmt.Println("Добавление базилика")
	}
	fmt.Println("Важно: тепловая обработка не требуется!")
}

// подача салата на тарелку
func (s *SaladPreparation) Plate() {
	fmt.Println("Шаг 4: Подача на тарелку")
	fmt.Println("Большая плоская тарелка")
	fmt.Println("Украшение зеленью")
}

// украшение салата
func (s *SaladPreparation) Garnish() {
	fmt.Println("Шаг 5: Украшение салата")
	if s.Type == "цезарь" {
		fmt.Println("Ломтики пармезана сверху")
	} else if s.Type == "греческий" {
		fmt.Println("Орегано и оливки")
	} else {
		fmt.Println("Листья базилика")
	}
}

// время приготовления салата
func (s *SaladPreparation) GetCookingTime() int {
	if s.Type == "цезарь" {
		return 20
	}
	return 15
}
