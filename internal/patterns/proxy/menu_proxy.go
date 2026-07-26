package proxy

import (
	"fmt"
	"time"
)

type RealMenu struct {
	items    []MenuItemProxy
	loadedAt time.Time
}

type MenuItemProxy struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

func NewRealMenu() *RealMenu {
	fmt.Println(" Загрузка меню из базы данных...")
	time.Sleep(1 * time.Second) // имитация загрузки

	return &RealMenu{
		items: []MenuItemProxy{
			{ID: 1, Name: "Пицца Маргарита", Price: 450, Category: "pizza"},
			{ID: 2, Name: "Паста Карбонара", Price: 380, Category: "pasta"},
			{ID: 3, Name: "Окономияки", Price: 500, Category: "pizza"},
			{ID: 4, Name: "Рамен", Price: 420, Category: "pasta"},
			{ID: 5, Name: "Салат Цезарь", Price: 280, Category: "salad"},
			{ID: 6, Name: "Тортилья", Price: 470, Category: "pizza"},
			{ID: 7, Name: "Чили кон карне", Price: 450, Category: "pasta"},
			{ID: 8, Name: "Лимончелло", Price: 150, Category: "drink"},
		},
		loadedAt: time.Now(),
	}
}

func (m *RealMenu) GetAllItems() []MenuItemProxy {
	return m.items
}

func (m *RealMenu) GetItemByID(id int) *MenuItemProxy {
	for _, item := range m.items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (m *RealMenu) GetItemsCount() int {
	return len(m.items)
}

// ПРОКСИ ДЛЯ МЕНЮ (КЭШИРОВАНИЕ)

type MenuProxy struct {
	realMenu    *RealMenu
	cachedItems []MenuItemProxy
	cacheTime   time.Time
	cacheTTL    time.Duration // время жизни кэша
}

func NewMenuProxy() *MenuProxy {
	return &MenuProxy{
		realMenu:    nil,
		cachedItems: []MenuItemProxy{},
		cacheTTL:    30 * time.Second,
	}
}

// загружает реальное меню (с кэшированием)
func (p *MenuProxy) GetRealMenu() *RealMenu {
	if p.realMenu != nil && time.Since(p.cacheTime) < p.cacheTTL {
		fmt.Println("Меню из кэша (актуально)")
		return p.realMenu
	}

	fmt.Println("Загрузка меню из базы данных...")
	p.realMenu = NewRealMenu()
	p.cacheTime = time.Now()
	fmt.Printf("Меню загружено (закешировано до %s)\n",
		p.cacheTime.Add(p.cacheTTL).Format("15:04:05"))

	return p.realMenu
}

func (p *MenuProxy) GetAllItems() []MenuItemProxy {
	return p.GetRealMenu().GetAllItems()
}

func (p *MenuProxy) GetItemByID(id int) *MenuItemProxy {
	return p.GetRealMenu().GetItemByID(id)
}

func (p *MenuProxy) GetItemsCount() int {
	return p.GetRealMenu().GetItemsCount()
}

// инвалидирует кэш
func (p *MenuProxy) InvalidateCache() {
	fmt.Println("Кэш меню инвалидирован")
	p.realMenu = nil
	p.cachedItems = []MenuItemProxy{}
}

// возвращает статус кэша
func (p *MenuProxy) GetCacheStatus() string {
	if p.realMenu == nil {
		return "Кэш пуст"
	}
	remaining := p.cacheTTL - time.Since(p.cacheTime)
	if remaining > 0 {
		return fmt.Sprintf("Кэш актуален (осталось %.0f сек)", remaining.Seconds())
	}
	return "Кэш устарел"
}
