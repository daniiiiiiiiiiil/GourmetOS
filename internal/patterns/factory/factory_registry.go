package factory

import "fmt"

type FactoryRegistry struct {
	factories map[string]DishFactory
}

func NewFactoryRegistry() *FactoryRegistry {
	return &FactoryRegistry{
		factories: make(map[string]DishFactory),
	}
}

func (r *FactoryRegistry) Register(name string, factory DishFactory) {
	r.factories[name] = factory
}

func (r *FactoryRegistry) GetFactory(name string) (DishFactory, error) { //1 шаг идет суда запрос например Японская кухня
	factory, ok := r.factories[name]
	if !ok {
		return nil, fmt.Errorf("кухня '%s' не найдена", name)
	}
	return factory, nil
}

func (r *FactoryRegistry) GetAllFactories() map[string]DishFactory {
	return r.factories
}

func (r *FactoryRegistry) GetAvailableKitchens() []string {
	kitchens := make([]string, 0, len(r.factories))
	for name := range r.factories {
		kitchens = append(kitchens, name)
	}
	return kitchens
}

type MenuItem struct {
	Category    string
	Dish        Dish
	KitchenName string
}

func (r *FactoryRegistry) CreateFullMenu() []MenuItem {
	var menu []MenuItem

	for kitchenName, factory := range r.factories {
		menu = append(menu, MenuItem{
			Category:    "Пицца",
			Dish:        factory.CreatePizza(),
			KitchenName: kitchenName,
		})
		menu = append(menu, MenuItem{
			Category:    "Паста",
			Dish:        factory.CreatePasta(),
			KitchenName: kitchenName,
		})
		menu = append(menu, MenuItem{
			Category:    "Салат",
			Dish:        factory.CreateSalad(),
			KitchenName: kitchenName,
		})
		menu = append(menu, MenuItem{
			Category:    "Напиток",
			Dish:        factory.CreateDrink(),
			KitchenName: kitchenName,
		})
	}

	return menu
}
