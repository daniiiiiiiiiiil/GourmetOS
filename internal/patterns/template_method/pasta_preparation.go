package template_method

import "fmt"

// приготовление пасты
type PastaPreparation struct {
	*DishPreparation
	Type      string // "карбонара", "болоньезе", "альфредо"
	PastaType string // "спагетти", "феттучини", "пенне"
}

func NewPastaPreparation(pastaType, sauceType string) *PastaPreparation {
	return &PastaPreparation{
		DishPreparation: NewDishPreparation("Паста " + sauceType),
		Type:            sauceType,
		PastaType:       pastaType,
	}
}

// сбор ингредиентов для пасты
func (p *PastaPreparation) GatherIngredients() {
	fmt.Printf("Шаг 1: Сбор ингредиентов для пасты %s\n", p.Type)
	fmt.Printf("Паста: %s\n", p.PastaType)
	fmt.Println("Оливковое масло, чеснок")
	if p.Type == "карбонара" {
		fmt.Println("Бекон, яйца, сыр Пармезан, перец")
	} else if p.Type == "болоньезе" {
		fmt.Println("Говяжий фарш, томаты, лук, морковь")
	} else {
		fmt.Println("Сливки, сыр, масло")
	}
}

// подготовка ингредиентов для пасты
func (p *PastaPreparation) PrepareIngredients() {
	fmt.Println("Шаг 2: Подготовка ингредиентов")
	fmt.Printf("Варка воды для пасты\n")
	fmt.Printf("Нарезка ингредиентов для соуса\n")
	if p.Type == "карбонара" {
		fmt.Println("Обжарка бекона до хруста")
	} else if p.Type == "болоньезе" {
		fmt.Println("Обжарка лука и моркови")
	}
}

// приготовление пасты (основной шаг)
func (p *PastaPreparation) Cook() {
	fmt.Println("Шаг 3: Приготовление пасты")
	fmt.Printf("Варка %s в подсоленной воде (8-10 мин)\n", p.PastaType)
	fmt.Printf("Приготовление соуса %s\n", p.Type)
	if p.Type == "карбонара" {
		fmt.Println("Смешивание яиц с пармезаном")
		fmt.Println("Соединение пасты с соусом")
	} else if p.Type == "болоньезе" {
		fmt.Println("Тушение соуса 30 минут")
		fmt.Println("Соединение пасты с соусом")
	} else {
		fmt.Println("Смешивание сливок с сыром")
		fmt.Println("Соединение пасты с соусом")
	}
}

// подача пасты на тарелку
func (p *PastaPreparation) Plate() {
	fmt.Println("Шаг 4: Подача на тарелку")
	fmt.Println("Глубокая тарелка")
	fmt.Println("Сверху пармезан и перец")
}

// украшение пасты
func (p *PastaPreparation) Garnish() {
	fmt.Println("Шаг 5: Украшение пасты")
	fmt.Println("Листики петрушки")
	fmt.Println("Тертый пармезан")
}

// время приготовления пасты
func (p *PastaPreparation) GetCookingTime() int {
	if p.Type == "болоньезе" {
		return 45
	}
	return 25
}
