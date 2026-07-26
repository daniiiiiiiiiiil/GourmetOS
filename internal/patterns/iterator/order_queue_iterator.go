package iterator

type OrderQueueItem struct {
	OrderID  int
	TableID  int
	Status   string
	Items    []string
	Total    float64
	Priority int // 1 - высокий, 2 - средний, 3 - низкий
}

type OrderQueueCollection struct {
	items []OrderQueueItem
}

// конструктор очереди заказов
func NewOrderQueueCollection() *OrderQueueCollection {
	return &OrderQueueCollection{
		items: []OrderQueueItem{},
	}
}

// добавляет заказ в очередь
func (q *OrderQueueCollection) Add(item OrderQueueItem) {
	q.items = append(q.items, item)
}

// создает итератор для очереди
func (q *OrderQueueCollection) CreateIterator() Iterator {
	return &OrderQueueIterator{
		collection: q,
		index:      0,
	}
}

// создает итератор с сортировкой по приоритету
func (q *OrderQueueCollection) CreatePriorityIterator() Iterator {
	sorted := &OrderQueueCollection{
		items: make([]OrderQueueItem, len(q.items)),
	}
	copy(sorted.items, q.items)

	// Пузырьковая сортировка по приоритету (1 - самый высокий)
	for i := 0; i < len(sorted.items)-1; i++ {
		for j := 0; j < len(sorted.items)-i-1; j++ {
			if sorted.items[j].Priority > sorted.items[j+1].Priority {
				sorted.items[j], sorted.items[j+1] = sorted.items[j+1], sorted.items[j]
			}
		}
	}

	return &OrderQueueIterator{
		collection: sorted,
		index:      0,
	}
}

// возвращает размер очереди
func (q *OrderQueueCollection) GetSize() int {
	return len(q.items)
}

// возвращает заказ по индексу
func (q *OrderQueueCollection) GetItem(index int) interface{} {
	if index < 0 || index >= len(q.items) {
		return nil
	}
	return q.items[index]
}

// возвращает активные заказы
func (q *OrderQueueCollection) GetActiveOrders() []OrderQueueItem {
	result := []OrderQueueItem{}
	for _, item := range q.items {
		if item.Status != "completed" && item.Status != "cancelled" {
			result = append(result, item)
		}
	}
	return result
}

// итератор для обхода очереди заказов
type OrderQueueIterator struct {
	collection *OrderQueueCollection
	index      int
}

// проверяет, есть ли следующий заказ
func (i *OrderQueueIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

// возвращает следующий заказ
func (i *OrderQueueIterator) Next() interface{} {
	if !i.HasNext() {
		return nil
	}
	item := i.collection.items[i.index]
	i.index++
	return item
}

// возвращает текущую позицию
func (i *OrderQueueIterator) GetIndex() int {
	return i.index
}
