package domain

import (
	"GourmetOS/pkg/errors"
	"time"
)

type Customer struct {
	CustomerID   int
	Name         string
	Phone        string
	Email        string
	Address      string
	BirthDate    time.Time
	LoyaltyLevel string
	TotalOrders  int
	TotalSpent   float64
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCustomer(CustomerID int,
	Name string,
	Phone string,
	Email string,
	Address string,
	BirthDate time.Time,
	LoyaltyLevel string,
	TotalOrders int,
	TotalSpent float64,
	IsActive bool,
	CreatedAt time.Time,
	UpdatedAt time.Time) *Customer {
	return &Customer{
		CustomerID:   CustomerID,
		Name:         Name,
		Phone:        Phone,
		Email:        Email,
		Address:      Address,
		BirthDate:    BirthDate,
		LoyaltyLevel: LoyaltyLevel,
		TotalOrders:  TotalOrders,
		TotalSpent:   TotalSpent,
		IsActive:     IsActive,
		CreatedAt:    CreatedAt,
		UpdatedAt:    UpdatedAt,
	}
}

// ещё надо будет проверки на длину сделать
func (c *Customer) Validate() error {
	var errs errors.ValidationErrors
	if c.Name == "" || len(c.Name) > 100 || len(c.Name) < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Name",
			Message: "Имя не может быть меньше 0 и больше 100 символов",
		})
	}
	if c.Phone == "" || len(c.Phone) < 0 || len(c.Phone) > 20 {
		errs = append(errs, errors.ValidationError{
			Field:   "Phone",
			Message: "Номер телефона не может быть меньше 0 или больше 20 символов",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
