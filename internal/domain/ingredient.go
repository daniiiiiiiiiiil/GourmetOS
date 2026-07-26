package domain

import (
	"GourmetOS/pkg/errors"
	"time"
)

type Ingredient struct {
	IngredientID  int
	Name          string
	Unit          string
	StockQuantity float64
	MinStock      float64
	MaxStock      float64
	CostPerUnit   float64
	Supplier      string
	ExpiryDate    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewIngredient(
	IngredientID int,
	Name string,
	Unit string,
	StockQuantity float64,
	MinStock float64,
	MaxStock float64,
	CostPerUnit float64,
	Supplier string,
	ExpiryDate time.Time,
	CreatedAt time.Time,
	UpdatedAt time.Time,
) *Ingredient {
	return &Ingredient{
		IngredientID:  IngredientID,
		Name:          Name,
		Unit:          Unit,
		StockQuantity: StockQuantity,
		MinStock:      MinStock,
		MaxStock:      MaxStock,
		CostPerUnit:   CostPerUnit,
		Supplier:      Supplier,
		ExpiryDate:    ExpiryDate,
		CreatedAt:     CreatedAt,
		UpdatedAt:     UpdatedAt,
	}
}

func (i *Ingredient) Validate() error {
	var errs errors.ValidationErrors
	if i.Name == "" || len(i.Name) == 0 || len(i.Name) > 100 {
		errs = append(errs, errors.ValidationError{
			Field:   "name",
			Message: "Имя не может быть пустым или больше 100 символов",
		})
	}
	if i.Unit == "" || len(i.Unit) == 0 || len(i.Unit) > 20 {
		errs = append(errs, errors.ValidationError{
			Field:   "unit",
			Message: "Unit не может быть пустым или больше 20 символов",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
