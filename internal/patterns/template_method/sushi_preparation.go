package template_method

import "fmt"

type SushiPreparation struct {
	*DishPreparation
	Type string // "нигири", "маки", "сашими"
	Fish string
}

func NewSushiPreparation(sushiType, fish string) *SushiPreparation {
	return &SushiPreparation{
		DishPreparation: NewDishPreparation("Суши " + sushiType),
		Type:            sushiType,
		Fish:            fish,
	}
}

// сбор ингредиентов для суши
func (s *SushiPreparation) GatherIngredients() {
	fmt.Printf("Шаг 1: Сбор ингредиентов для суши %s\n", s.Type)
	fmt.Println("Рис для суши, рисовый уксус")
	fmt.Printf("Свежая рыба: %s\n", s.Fish)
	fmt.Println("Нори, васаби, имбирь")
	fmt.Println("Соевый соус")
}

// подготовка ингредиентов для суши
func (s *SushiPreparation) PrepareIngredients() {
	fmt.Println("Шаг 2: Подготовка ингредиентов")
	fmt.Println("Варка риса для суши (важный этап)")
	fmt.Println("Заправка риса уксусом")
	fmt.Printf("Нарезка рыбы на тонкие ломтики\n")
	if s.Type == "маки" {
		fmt.Println("Подготовка листов нори")
	}
}

// приготовление суши (особенность: не требует тепловой обработки)
func (s *SushiPreparation) Cook() {
	fmt.Println("Шаг 3: Формирование суши")
	if s.Type == "нигири" {
		fmt.Println("Формирование рисовых подушечек")
		fmt.Printf("Укладывание %s сверху\n", s.Fish)
	} else if s.Type == "маки" {
		fmt.Println("Раскладывание риса на нори")
		fmt.Printf("Добавление %s и овощей\n", s.Fish)
		fmt.Println("Скручивание рулета")
		fmt.Println("Нарезка на 6 кусков")
	} else {
		fmt.Printf("Тонкая нарезка %s\n", s.Fish)
		fmt.Println("Красивая раскладка на тарелке")
	}
	fmt.Println("Важно: тепловая обработка не требуется!")
}

// подача суши на тарелку
func (s *SushiPreparation) Plate() {
	fmt.Println("Шаг 4: Подача на тарелку")
	fmt.Println("Керамическая тарелка")
	fmt.Println("Рядом васаби и маринованный имбирь")
}

// украшение суши
func (s *SushiPreparation) Garnish() {
	fmt.Println("Шаг 5: Украшение суши")
	fmt.Println("Листья шисо (японский базилик)")
	fmt.Println("Семена кунжута")
}

// время приготовления суши
func (s *SushiPreparation) GetCookingTime() int {
	return 30
}
