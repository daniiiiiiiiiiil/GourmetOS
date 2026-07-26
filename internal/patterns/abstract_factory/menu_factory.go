package abstract_factory

// интерфейс абстрактной фабрики
// Каждая фабрика создает семейство связанных продуктов
type MenuFactory interface {
	CreateMainDish() *MainDish
	CreateSideDish() *SideDish
	CreateDrink() *Drink
	CreateDessert() *Dessert
	CreateMenuSet() *MenuSet
	GetName() string
	GetDescription() string
}
