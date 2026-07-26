package iterator

type MenuItem struct {
	Name        string
	Price       float64
	Category    string
	Cuisine     string
	Description string
}

type MenuCollection struct {
	items []MenuItem
}

func NewMenuCollection() *MenuCollection {
	return &MenuCollection{
		items: []MenuItem{},
	}
}

// добавляет блюдо в меню
func (m *MenuCollection) Add(item MenuItem) {
	m.items = append(m.items, item)
}

// создает итератор для меню
func (m *MenuCollection) CreateIterator() Iterator {
	return &MenuIterator{
		collection: m,
		index:      0,
	}
}

// возвращает количество блюд в меню
func (m *MenuCollection) GetSize() int {
	return len(m.items)
}

// GetItem — возвращает блюдо по индексу
func (m *MenuCollection) GetItem(index int) interface{} {
	if index < 0 || index >= len(m.items) {
		return nil
	}
	return m.items[index]
}

// возвращает все блюда
func (m *MenuCollection) GetAllItems() []MenuItem {
	return m.items
}

// фильтрует меню по категории
func (m *MenuCollection) FilterByCategory(category string) []MenuItem {
	result := []MenuItem{}
	for _, item := range m.items {
		if item.Category == category {
			result = append(result, item)
		}
	}
	return result
}

// итератор для обхода меню
type MenuIterator struct {
	collection *MenuCollection
	index      int
}

// проверяет, есть ли следующее блюдо
func (i *MenuIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

// возвращает следующее блюдо
func (i *MenuIterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	item := i.collection.items[i.index]
	i.index++
	return item
}

// возвращает текущую позицию
func (i *MenuIterator) GetIndex() int {
	return i.index
}
