package adapter

import (
	"fmt"
	"math/rand"
	"time"
)

type DeliveryAPI struct{}

func NewDeliveryAPI() *DeliveryAPI {
	return &DeliveryAPI{}
}

// создание заказа на доставку
func (d *DeliveryAPI) CreateOrder(address string, items []string, total float64) (string, error) {
	fmt.Printf("   [Delivery API] Создание заказа доставки\n")
	fmt.Printf("   Адрес: %s\n", address)
	fmt.Printf("   Блюда: %v\n", items)
	fmt.Printf("   Сумма: %.2f руб.\n", total)

	time.Sleep(500 * time.Millisecond)

	if rand.Float64() < 0.05 {
		return "", fmt.Errorf("Delivery API: недоступные курьеры")
	}

	orderID := fmt.Sprintf("DELIVERY_%d", time.Now().UnixNano())
	fmt.Printf("   [Delivery API] Заказ создан: %s\n", orderID)
	return orderID, nil
}

// статус доставки
func (d *DeliveryAPI) GetDeliveryStatus(orderID string) (string, error) {
	statuses := []string{"pending", "preparing", "on_way", "delivered"}
	status := statuses[rand.Intn(len(statuses))]
	fmt.Printf("   [Delivery API] Статус доставки %s: %s\n", orderID, status)
	return status, nil
}

// отмена доставки
func (d *DeliveryAPI) CancelDelivery(orderID string) error {
	fmt.Printf("   [Delivery API] Отмена доставки: %s\n", orderID)
	return nil
}

// адаптер для службы доставки
type DeliveryAdapter struct {
	api     *DeliveryAPI
	address string
}

// конструктор
func NewDeliveryAdapter(address string) *DeliveryAdapter {
	return &DeliveryAdapter{
		api:     NewDeliveryAPI(),
		address: address,
	}
}

// создает доставку
func (d *DeliveryAdapter) CreateDelivery(items []string, total float64) (string, error) {
	fmt.Printf("Создание доставки через Delivery Adapter\n")
	return d.api.CreateOrder(d.address, items, total)
}

func (d *DeliveryAdapter) GetDeliveryStatus(orderID string) (string, error) {
	fmt.Printf("🔍 Проверка статуса доставки: %s\n", orderID)
	return d.api.GetDeliveryStatus(orderID)
}

func (d *DeliveryAdapter) CancelDelivery(orderID string) error {
	fmt.Printf("Отмена доставки: %s\n", orderID)
	return d.api.CancelDelivery(orderID)
}

func (d *DeliveryAdapter) GetAdapterName() string {
	return "Delivery Adapter"
}
