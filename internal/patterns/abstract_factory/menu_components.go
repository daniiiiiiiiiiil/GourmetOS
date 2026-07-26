package abstract_factory

// ПРОДУКТЫ (блюда)

// основное блюдо
type MainDish struct {
	Name        string
	Description string
	Price       float64
	Weight      int
	CookingTime int
}

func (m *MainDish) GetInfo() string {
	return m.Name + " (" + m.Description + ") - " + string(rune(m.Price)) + " руб."
}

// гарнир
type SideDish struct {
	Name        string
	Description string
	Price       float64
	Weight      int
}

func (s *SideDish) GetInfo() string {
	return s.Name + " (" + s.Description + ") - " + string(rune(s.Price)) + " руб."
}

// напиток
type Drink struct {
	Name        string
	Description string
	Price       float64
	Volume      int
}

func (d *Drink) GetInfo() string {
	return d.Name + " (" + d.Description + ") - " + string(rune(d.Price)) + " руб."
}

// десерт
type Dessert struct {
	Name        string
	Description string
	Price       float64
	Weight      int
	Calories    int
}

func (d *Dessert) GetInfo() string {
	return d.Name + " (" + d.Description + ") - " + string(rune(d.Price)) + " руб."
}

// МЕНЮ (комплексный набор)

// комплексное меню (набор продуктов)
type MenuSet struct {
	Name        string
	Description string
	MainDish    *MainDish
	SideDish    *SideDish
	Drink       *Drink
	Dessert     *Dessert
	TotalPrice  float64
	Discount    float64
}

func NewMenuSet(name, description string, main *MainDish, side *SideDish, drink *Drink, dessert *Dessert, discount float64) *MenuSet {
	total := main.Price + side.Price + drink.Price + dessert.Price
	if discount > 0 {
		total = total - (total * discount / 100)
	}

	return &MenuSet{
		Name:        name,
		Description: description,
		MainDish:    main,
		SideDish:    side,
		Drink:       drink,
		Dessert:     dessert,
		TotalPrice:  total,
		Discount:    discount,
	}
}

func (m *MenuSet) GetTotalPrice() float64 {
	return m.TotalPrice
}

func (m *MenuSet) GetOriginalPrice() float64 {
	return m.MainDish.Price + m.SideDish.Price + m.Drink.Price + m.Dessert.Price
}

func (m *MenuSet) GetItems() []string {
	return []string{
		m.MainDish.Name,
		m.SideDish.Name,
		m.Drink.Name,
		m.Dessert.Name,
	}
}

func (m *MenuSet) GetInfo() string {
	info := " " + m.Name + "\n"
	info += "   " + m.Description + "\n"
	info += "   Основное: " + m.MainDish.Name + " (" + string(rune(m.MainDish.Price)) + " руб.)\n"
	info += "   Гарнир: " + m.SideDish.Name + " (" + string(rune(m.SideDish.Price)) + " руб.)\n"
	info += "   Напиток: " + m.Drink.Name + " (" + string(rune(m.Drink.Price)) + " руб.)\n"
	info += "   Десерт: " + m.Dessert.Name + " (" + string(rune(m.Dessert.Price)) + " руб.)\n"
	if m.Discount > 0 {
		info += "   Скидка: " + string(rune(m.Discount)) + "%\n"
		info += "   Итоговая цена: " + string(rune(m.TotalPrice)) + " руб. (экономия " + string(rune(m.GetOriginalPrice()-m.TotalPrice)) + " руб.)"
	} else {
		info += "   Цена: " + string(rune(m.TotalPrice)) + " руб."
	}
	return info
}
