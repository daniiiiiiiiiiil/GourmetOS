package domain

import (
	"GourmetOS/pkg/errors"
	"time"
)

type Employees struct {
	EmployeeID int       `json:"employee_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Role       string    `json:"role"`
	Shift      string    `json:"shift"`
	HireDate   time.Time `json:"hire_date"`
	Salary     float64   `json:"salary"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewEmployees(
	EmployeeID int,
	Name string,
	Email string,
	Phone string,
	Role string,
	Shift string,
	HireDate time.Time,
	Salary float64,
	IsActive bool,
	CreatedAt time.Time,
	UpdatedAt time.Time) *Employees {
	return &Employees{
		EmployeeID: EmployeeID,
		Name:       Name,
		Email:      Email,
		Phone:      Phone,
		Role:       Role,
		Shift:      Shift,
		HireDate:   HireDate,
		Salary:     Salary,
		IsActive:   IsActive,
		CreatedAt:  CreatedAt,
		UpdatedAt:  UpdatedAt,
	}
}

func (e *Employees) Validate() error {
	var errs errors.ValidationErrors
	if e.Name == "" || len(e.Name) > 100 || len(e.Name) < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Name",
			Message: "Имя не может быть меньше 0 и больше 100 символов",
		})
	}
	if e.Phone == "" || len(e.Phone) < 0 || len(e.Phone) > 20 {
		errs = append(errs, errors.ValidationError{
			Field:   "Phone",
			Message: "Номер телефона не может быть меньше 0 или больше 20 символов",
		})
	}
	//для емайла сделать лучше вадилацию поотм
	if e.Email == "" || len(e.Email) < 5 || len(e.Email) > 100 {
		errs = append(errs, errors.ValidationError{
			Field:   "Email",
			Message: "Email не может быть пустым и не может быть больше 100 символов",
		})
	}
	if e.Role == "" || len(e.Role) < 0 || len(e.Role) > 50 {
		errs = append(errs, errors.ValidationError{
			Field:   "Role",
			Message: "Роль не может быть пустой и не может быть больше 50 символов",
		})
	}
	if e.HireDate.IsZero() {
		errs = append(errs, errors.ValidationError{
			Field:   "HireDate",
			Message: "Дата не может быть пустой",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
