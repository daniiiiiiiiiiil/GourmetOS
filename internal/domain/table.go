package domain

import (
	"GourmetOS/pkg/errors"
	"time"
)

type Table struct {
	TableID    int
	Number     int
	Capacity   int
	Location   string
	IsOccupied bool
	IsReserved bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTable(
	tableID int,
	number int,
	capacity int,
	location string,
	isOccupied bool,
	isReserved bool,
	createdAt time.Time,
	updatedAt time.Time) *Table {
	return &Table{
		TableID:    tableID,
		Number:     number,
		Capacity:   capacity,
		Location:   location,
		IsOccupied: isOccupied,
		IsReserved: isReserved,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

func (t *Table) Validate() error {
	var errs errors.ValidationErrors
	if t.Number < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Number",
			Message: "Число не может быть меньше нуля",
		})
	}
	if t.Capacity < 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "Capacity",
			Message: "Вместимость не может быть меньше 0",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}
