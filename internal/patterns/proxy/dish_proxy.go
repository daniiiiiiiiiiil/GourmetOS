package proxy

import "fmt"

type RealDish struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    string
	Cuisine     string
	Ingredients []string
	ImageData   []byte
	isLoaded    bool
}

func NewRealDish(id int, name, description string, price float64, category, cuisine string, ingredients []string) *RealDish {
	return &RealDish{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Cuisine:     cuisine,
		Ingredients: ingredients,
		ImageData:   []byte{},
		isLoaded:    false,
	}
}

// загружает изображение (тяжелая операция)
func (d *RealDish) LoadImage() {
	if !d.isLoaded {
		fmt.Printf("Загрузка изображения для блюда '%s'...\n", d.Name)
		// Имитация загрузки больших данных
		d.ImageData = []byte(fmt.Sprintf("image_data_%d", d.ID))
		d.isLoaded = true
		fmt.Printf("Изображение загружено (%d байт)\n", len(d.ImageData))
	}
}

// возвращает информацию о блюде
func (d *RealDish) GetInfo() string {
	return fmt.Sprintf("%s (%.2f руб.) - %s", d.Name, d.Price, d.Description)
}

// ПРОКСИ ДЛЯ БЛЮДА

type DishProxy struct {
	id          int
	name        string
	description string
	price       float64
	category    string
	cuisine     string
	ingredients []string
	realDish    *RealDish
}

func NewDishProxy(id int, name, description string, price float64, category, cuisine string, ingredients []string) *DishProxy {
	return &DishProxy{
		id:          id,
		name:        name,
		description: description,
		price:       price,
		category:    category,
		cuisine:     cuisine,
		ingredients: ingredients,
		realDish:    nil,
	}
}

// загружает реальный объект (ленивая инициализация)
func (p *DishProxy) GetRealDish() *RealDish {
	if p.realDish == nil {
		fmt.Printf("Инициализация реального блюда '%s'...\n", p.name)
		p.realDish = NewRealDish(
			p.id, p.name, p.description,
			p.price, p.category, p.cuisine,
			p.ingredients,
		)
	}
	return p.realDish
}

// возвращает имя (без загрузки реального объекта)
func (p *DishProxy) GetName() string {
	return p.name
}

// возвращает цену (без загрузки реального объекта)
func (p *DishProxy) GetPrice() float64 {
	return p.price
}

// возвращает информацию (загружает реальный объект только если нужно)
func (p *DishProxy) GetInfo() string {
	return p.GetRealDish().GetInfo()
}

// загружает изображение (делегирует реальному объекту)
func (p *DishProxy) LoadImage() {
	p.GetRealDish().LoadImage()
}

// возвращает категорию
func (p *DishProxy) GetCategory() string {
	return p.category
}

// проверяет, загружен ли реальный объект
func (p *DishProxy) IsLoaded() bool {
	return p.realDish != nil
}
