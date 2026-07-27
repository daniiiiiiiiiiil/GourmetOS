package domain

import (
	"GourmetOS/pkg/errors"
	"time"
)

type Dish struct {
	DishID       int
	Name         string
	Description  string
	Price        float64
	Category     string
	Cuisine      string
	CookingTime  int
	IsAvailable  bool
	IsVegetarian bool
	IsVegan      bool
	IsGlutenFree bool
	Calories     *int
	ImageURL     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewDish(dishID int,
	name string,
	description string,
	price float64,
	category string,
	cuisine string,
	cookingTime int,
	isAvailable bool,
	isVegetarian bool,
	isVegan bool,
	isGlutenFree bool,
	calories *int,
	imageURL *string,
	createdAt time.Time,
	updatedAt time.Time) *Dish {
	return &Dish{
		DishID:       dishID,
		Name:         name,
		Description:  description,
		Price:        price,
		Category:     category,
		Cuisine:      cuisine,
		CookingTime:  cookingTime,
		IsAvailable:  isAvailable,
		IsVegetarian: isVegetarian,
		IsVegan:      isVegan,
		IsGlutenFree: isGlutenFree,
		Calories:     calories,
		ImageURL:     imageURL,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

// ещё надо будет проверки на длину сделать
func (dish *Dish) Validate() error {
	var errs errors.ValidationErrors
	if dish.Name == "" || len(dish.Name) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "name",
			Message: "Имя не может быть пустым",
		})
	}
	if dish.Price < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "price",
			Message: "Цена не может быть меньше 0",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
