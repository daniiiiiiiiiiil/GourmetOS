package dto

import (
	"GourmetOS/internal/domain"
	"GourmetOS/pkg/errors"
	"GourmetOS/pkg/pagination"
	"time"
)

type CreateDishRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Cuisine      string  `json:"cuisine"`
	CookingTime  int     `json:"cooking_time"`
	IsVegetarian bool    `json:"is_vegetarian"`
	IsVegan      bool    `json:"is_vegan"`
	IsGlutenFree bool    `json:"is_gluten_free"`
	Calories     *int    `json:"calories"`
	ImageURL     *string `json:"image_url"`
}

func (r *CreateDishRequest) Validate() error {
	var errs errors.ValidationErrors
	if r.Name == "" || len(r.Name) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Name",
			Message: "Имя не может быть пустым",
		})
	}
	if r.Price <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Price",
			Message: "Цена не может быть меньше или равен нулю",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

func (r *CreateDishRequest) ToDomain() domain.Dish {
	return domain.Dish{
		Name:         r.Name,
		Description:  r.Description,
		Price:        r.Price,
		Category:     r.Category,
		CookingTime:  r.CookingTime,
		IsVegetarian: r.IsVegetarian,
		IsVegan:      r.IsVegan,
		IsGlutenFree: r.IsGlutenFree,
		Calories:     r.Calories,
		ImageURL:     r.ImageURL,
	}
}

type UpdateDishRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Cuisine      string  `json:"cuisine"`
	CookingTime  int     `json:"cooking_time"`
	IsVegetarian bool    `json:"is_vegetarian"`
	IsVegan      bool    `json:"is_vegan"`
	IsGlutenFree bool    `json:"is_gluten_free"`
	Calories     *int    `json:"calories"`
	ImageURL     *string `json:"image_url"`
}

func (r *UpdateDishRequest) Validate() error {
	var errs errors.ValidationErrors
	if r.Name == "" || len(r.Name) == 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Name",
			Message: "Имя не может быть пустым",
		})
	}
	if r.Price <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Price",
			Message: "Цена не может быть меньше или равен нулю",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

func (r *UpdateDishRequest) ToDomain() domain.Dish {
	return domain.Dish{
		Name:         r.Name,
		Description:  r.Description,
		Price:        r.Price,
		Category:     r.Category,
		CookingTime:  r.CookingTime,
		IsVegetarian: r.IsVegetarian,
		IsVegan:      r.IsVegan,
		IsGlutenFree: r.IsGlutenFree,
		Calories:     r.Calories,
		ImageURL:     r.ImageURL,
	}
}

type DishResponse struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Price        float64    `json:"price"`
	Category     string     `json:"category"`
	Cuisine      string     `json:"cuisine"`
	CookingTime  int        `json:"cooking_time"`
	IsAvailable  bool       `json:"is_available"`
	IsVegetarian bool       `json:"is_vegetarian"`
	IsVegan      bool       `json:"is_vegan"`
	IsGlutenFree bool       `json:"is_gluten_free"`
	Calories     *int       `json:"calories"`
	ImageURL     *string    `json:"image_url"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func DishResponseFromDomain(dish domain.Dish) DishResponse {
	return DishResponse{
		ID:           dish.DishID,
		Name:         dish.Name,
		Description:  dish.Description,
		Price:        dish.Price,
		Category:     dish.Category,
		Cuisine:      dish.Cuisine,
		CookingTime:  dish.CookingTime,
		IsAvailable:  dish.IsAvailable,
		IsVegetarian: dish.IsVegetarian,
		IsVegan:      dish.IsVegan,
		IsGlutenFree: dish.IsGlutenFree,
		Calories:     dish.Calories,
		ImageURL:     dish.ImageURL,
		CreatedAt:    dish.CreatedAt,
		UpdatedAt:    &dish.UpdatedAt,
	}
}

type DishListResponse struct {
	Dishes     []DishResponse        `json:"dishes"`
	Pagination pagination.Pagination `json:"pagination"`
}

func NewDishListResponse(dishes []domain.Dish, total, limit, offset int) DishListResponse {
	resp := DishListResponse{
		Dishes:     make([]DishResponse, 0, len(dishes)),
		Pagination: pagination.NewPagination(total, limit, offset),
	}
	for _, dish := range dishes {
		resp.Dishes = append(resp.Dishes, DishResponseFromDomain(dish))
	}
	return resp
}

type UpdateAvailabilityRequest struct {
	IsAvailable bool `json:"is_available" binding:"required"`
}
