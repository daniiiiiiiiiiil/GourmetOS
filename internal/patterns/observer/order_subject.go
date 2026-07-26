package observer

type OrderSubject struct {
	observers []Observer // Список подписчиков
}

func NewOrderSubject() *OrderSubject {
	return &OrderSubject{
		observers: []Observer{},
	}
}

// подписать наблюдателя
func (s *OrderSubject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

// отписать наблюдателя
func (s *OrderSubject) RemoveObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// оповестить всех подписчиков
func (s *OrderSubject) NotifyObservers(event Event) {
	for _, observer := range s.observers {
		observer.Update(event)
	}
}

func (s *OrderSubject) GetObservers() []Observer {
	return s.observers
}

func (s *OrderSubject) GetObserversCount() int {
	return len(s.observers)
}
