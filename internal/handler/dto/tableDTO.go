package dto

import (
	"GourmetOS/internal/domain"
	"GourmetOS/pkg/errors"
	"GourmetOS/pkg/pagination"
	"time"
)

type CreateTableRequest struct {
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`
	Location string `json:"location"`
}

func (c *CreateTableRequest) Validate() error {
	var errs errors.ValidationErrors
	if c.Number <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "number",
			Message: "Номер не может быть меньше или равен 0",
		})
	}
	if c.Capacity <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "capacity",
			Message: "Вместимость не может быть меньше или равен 0",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

func (c *CreateTableRequest) ToDomain() *domain.Table {
	return &domain.Table{
		Number:   c.Number,
		Capacity: c.Capacity,
		Location: c.Location,
	}
}

type UpdateTableRequest struct {
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`
	Location string `json:"location"`
}

func (c *UpdateTableRequest) Validate() error {
	var errs errors.ValidationErrors
	if c.Number <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "number",
			Message: "Номер не может быть меньше или равен 0",
		})
	}
	if c.Capacity <= 0 {
		errs = append(errs, errors.ValidationError{
			Field:   "capacity",
			Message: "Вместимость не может быть меньше или равен 0",
		})
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

func (c *UpdateTableRequest) ToDomain() *UpdateTableRequest {
	return &UpdateTableRequest{
		Number:   c.Number,
		Capacity: c.Capacity,
		Location: c.Location,
	}
}

type TableResponse struct {
	TableID    int       `json:"table_id"`
	Number     int       `json:"number"`
	Capacity   int       `json:"capacity"`
	Location   string    `json:"location"`
	IsOccupied bool      `json:"is_occupied"`
	IsReserved bool      `json:"is_reserved"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewTableResponseFromDomain(table domain.Table) *TableResponse {
	return &TableResponse{
		TableID:    table.TableID,
		Number:     table.Number,
		Capacity:   table.Capacity,
		Location:   table.Location,
		IsOccupied: table.IsOccupied,
		IsReserved: table.IsReserved,
		CreatedAt:  table.CreatedAt,
		UpdatedAt:  table.UpdatedAt,
	}
}

type TableListResponse struct {
	TableList  []TableResponse       `json:"table_list"`
	Pagination pagination.Pagination `json:"pagination"`
}

func NewTableListResponse(tables []domain.Table, total, limit, offset int) *TableListResponse {
	resp := &TableListResponse{
		TableList:  make([]TableResponse, 0, len(tables)),
		Pagination: pagination.NewPagination(total, limit, offset),
	}
	for _, table := range tables {
		resp.TableList = append(resp.TableList, *NewTableResponseFromDomain(table))
	}
	return resp
}

func FromDomainTableResponse(t domain.Table) *TableResponse {
	return &TableResponse{
		TableID:    t.TableID,
		Number:     t.Number,
		Capacity:   t.Capacity,
		Location:   t.Location,
		IsOccupied: t.IsOccupied,
		IsReserved: t.IsReserved,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}
