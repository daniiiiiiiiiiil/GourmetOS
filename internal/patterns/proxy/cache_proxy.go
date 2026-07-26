package proxy

import (
	"fmt"
	"sync"
	"time"
)

// интерфейс для объектов, которые можно кэшировать
type Cacheable interface {
	GetID() int
	GetData() string
}

// общий прокси для кэширования
type CacheProxy struct {
	cache  map[int]CacheItem
	mu     sync.RWMutex
	ttl    time.Duration
	loader func(id int) (Cacheable, error)
}

// элемент кэша
type CacheItem struct {
	data     Cacheable
	loadedAt time.Time
}

func NewCacheProxy(ttl time.Duration, loader func(id int) (Cacheable, error)) *CacheProxy {
	return &CacheProxy{
		cache:  make(map[int]CacheItem),
		ttl:    ttl,
		loader: loader,
	}
}

func (p *CacheProxy) Get(id int) (Cacheable, error) {
	p.mu.RLock()
	item, exists := p.cache[id]
	p.mu.RUnlock()

	if exists && time.Since(item.loadedAt) < p.ttl {
		fmt.Printf("Объект #%d из кэша\n", id)
		return item.data, nil
	}

	fmt.Printf("Загрузка объекта #%d из источника\n", id)
	data, err := p.loader(id)
	if err != nil {
		return nil, err
	}

	// Сохраняем в кэш
	p.mu.Lock()
	p.cache[id] = CacheItem{
		data:     data,
		loadedAt: time.Now(),
	}
	p.mu.Unlock()

	fmt.Printf("Объект #%d закеширован\n", id)
	return data, nil
}

func (p *CacheProxy) Invalidate(id int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.cache, id)
	fmt.Printf("Объект #%d удалён из кэша\n", id)
}

func (p *CacheProxy) Clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cache = make(map[int]CacheItem)
	fmt.Println("Весь кэш очищен")
}

func (p *CacheProxy) GetStats() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return fmt.Sprintf("Кэш: %d объектов", len(p.cache))
}

// ПРИМЕР ИСПОЛЬЗОВАНИЯ: КЭШИРОВАНИЕ КЛИЕНТОВ

// клиент для кэширования
type CachedCustomer struct {
	ID    int
	Name  string
	Phone string
	Email string
}

func (c *CachedCustomer) GetID() int {
	return c.ID
}

func (c *CachedCustomer) GetData() string {
	return fmt.Sprintf("Клиент: %s (%s)", c.Name, c.Phone)
}

// загрузчик клиентов
func CustomerLoader(id int) (Cacheable, error) {
	time.Sleep(200 * time.Millisecond)
	customers := map[int]CachedCustomer{
		1: {ID: 1, Name: "Андрей Морозов", Phone: "+7-999-111-22-33", Email: "andrey@mail.com"},
		2: {ID: 2, Name: "Екатерина Новикова", Phone: "+7-999-222-33-44", Email: "ekaterina@mail.com"},
		3: {ID: 3, Name: "Михаил Федоров", Phone: "+7-999-333-44-55", Email: "mikhail@mail.com"},
	}
	if customer, exists := customers[id]; exists {
		return &customer, nil
	}
	return nil, fmt.Errorf("клиент #%d не найден", id)
}
