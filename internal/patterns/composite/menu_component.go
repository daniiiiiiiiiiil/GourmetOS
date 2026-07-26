package composite

// интерфейс для всех компонентов меню
// Может быть как листом (блюдо), так и композитом (категория, комбо)
type MenuComponent interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
	Print(indent string)
	Add(child MenuComponent)
	Remove(child MenuComponent)
	GetChild(index int) MenuComponent
	IsComposite() bool
}

//	базовая реализация для всех компонентов
//
// Используется для избежания дублирования кода
type BaseComponent struct {
	name        string
	price       float64
	description string
}

func (b *BaseComponent) GetName() string {
	return b.name
}

func (b *BaseComponent) GetPrice() float64 {
	return b.price
}

func (b *BaseComponent) GetDescription() string {
	return b.description
}

func (b *BaseComponent) Print(indent string) {
	// Будет переопределено в наследниках
}

func (b *BaseComponent) Add(child MenuComponent) {
	// По умолчанию ничего не делает (для листьев)
}

func (b *BaseComponent) Remove(child MenuComponent) {
	// По умолчанию ничего не делает (для листьев)
}

func (b *BaseComponent) GetChild(index int) MenuComponent {
	return nil
}

func (b *BaseComponent) IsComposite() bool {
	return false
}
