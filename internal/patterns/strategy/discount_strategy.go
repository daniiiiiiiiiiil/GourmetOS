package strategy

import "fmt"

type DiscountStrategy interface {
	ApplyDiscount(amount float64) float64
	GetName() string
}

type NoDiscountStrategy struct{}

func NewNoDiscountStrategy() *NoDiscountStrategy {
	return &NoDiscountStrategy{}
}

func (d *NoDiscountStrategy) ApplyDiscount(amount float64) float64 {
	return amount
}

func (d *NoDiscountStrategy) GetName() string {
	return "Без скидки"
}

type SeasonalDiscountStrategy struct{}

func NewSeasonalDiscountStrategy() *SeasonalDiscountStrategy {
	return &SeasonalDiscountStrategy{}
}

func (d *SeasonalDiscountStrategy) ApplyDiscount(amount float64) float64 {
	discount := amount * 0.10
	fmt.Printf(" Сезонная скидка: %.2f руб. (10%%)\n", discount)
	return amount - discount
}

func (d *SeasonalDiscountStrategy) GetName() string {
	return "Сезонная скидка (10%)"
}

type LoyaltyDiscountStrategy struct {
	Visits int
}

func NewLoyaltyDiscountStrategy(visits int) *LoyaltyDiscountStrategy {
	return &LoyaltyDiscountStrategy{Visits: visits}
}

func (d *LoyaltyDiscountStrategy) ApplyDiscount(amount float64) float64 {
	var percent float64
	switch {
	case d.Visits >= 50:
		percent = 0.20
	case d.Visits >= 20:
		percent = 0.15
	case d.Visits >= 10:
		percent = 0.10
	default:
		percent = 0.05
	}

	discount := amount * percent
	fmt.Printf("Скидка постоянного клиента: %.2f руб. (%.0f%%)\n", discount, percent*100)
	return amount - discount
}

func (d *LoyaltyDiscountStrategy) GetName() string {
	return "Скидка постоянного клиента"
}

type BirthdayDiscountStrategy struct{}

func NewBirthdayDiscountStrategy() *BirthdayDiscountStrategy {
	return &BirthdayDiscountStrategy{}
}

func (d *BirthdayDiscountStrategy) ApplyDiscount(amount float64) float64 {
	discount := amount * 0.20
	fmt.Printf("Скидка в день рождения: %.2f руб. (20%%)\n", discount)
	fmt.Println("С Днём Рождения!")
	return amount - discount
}

func (d *BirthdayDiscountStrategy) GetName() string {
	return "Скидка в день рождения (20%)"
}

type OrderDiscountStrategy struct {
	MinAmount float64
	Discount  float64
}

func NewOrderDiscountStrategy(minAmount, discount float64) *OrderDiscountStrategy {
	return &OrderDiscountStrategy{
		MinAmount: minAmount,
		Discount:  discount,
	}
}

func (d *OrderDiscountStrategy) ApplyDiscount(amount float64) float64 {
	if amount >= d.MinAmount {
		fmt.Printf("Скидка на крупный заказ: %.2f руб.\n", d.Discount)
		return amount - d.Discount
	}
	return amount
}

func (d *OrderDiscountStrategy) GetName() string {
	return fmt.Sprintf("Скидка на заказ от %.0f руб.", d.MinAmount)
}
