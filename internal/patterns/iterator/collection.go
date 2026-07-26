package iterator

type Iterator interface {
	HasNext() bool     // есть ли следующий элемент
	Next() interface{} // возвращает следующий элемент
	GetIndex() int     // возвращает текущую позицию
}

type Collection interface {
	CreateIterator() Iterator      // создает итератор для этой коллекции
	GetSize() int                  // возвращает размер коллекции
	GetItem(index int) interface{} // возвращает элемент по индексу
}
