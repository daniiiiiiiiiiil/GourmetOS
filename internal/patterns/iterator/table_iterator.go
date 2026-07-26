package iterator

type Table struct {
	ID         int
	Number     int
	Capacity   int
	Location   string
	IsOccupied bool
}

type TableCollection struct {
	items []Table
}

func NewTableCollection() *TableCollection {
	return &TableCollection{
		items: []Table{},
	}
}

// добавляет стол
func (t *TableCollection) Add(item Table) {
	t.items = append(t.items, item)
}

// создает итератор для столов
func (t *TableCollection) CreateIterator() Iterator {
	return &TableIterator{
		collection: t,
		index:      0,
	}
}

// возвращает количество столов
func (t *TableCollection) GetSize() int {
	return len(t.items)
}

// возвращает стол по индексу
func (t *TableCollection) GetItem(index int) interface{} {
	if index < 0 || index >= len(t.items) {
		return nil
	}
	return t.items[index]
}

// возвращает свободные столы
func (t *TableCollection) GetFreeTables() []Table {
	result := []Table{}
	for _, table := range t.items {
		if !table.IsOccupied {
			result = append(result, table)
		}
	}
	return result
}

// итератор для обхода столов
type TableIterator struct {
	collection *TableCollection
	index      int
}

// проверяет, есть ли следующий стол
func (i *TableIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

// возвращает следующий стол
func (i *TableIterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	item := i.collection.items[i.index]
	i.index++
	return item
}

// возвращает текущую позицию
func (i *TableIterator) GetIndex() int {
	return i.index
}
