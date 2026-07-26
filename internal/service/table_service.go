package service

import (
	"GourmetOS/internal/domain"
	"GourmetOS/internal/repository/interfaceRepository"
	"GourmetOS/pkg/errors"
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type TableService struct {
	tableService interfaceRepository.TableRepository
	logger       *zap.Logger
}

func NewTableService(tableService interfaceRepository.TableRepository, logger *zap.Logger) *TableService {
	return &TableService{
		tableService: tableService,
		logger:       logger,
	}
}

func (t *TableService) CreateTableService(ctx context.Context, conn *pgx.Conn, table *domain.Table) (*domain.Table, error) {
	if err := table.Validate(); err != nil {
		return nil, errors.ValidationError{
			Field:   "table",
			Message: "Ошибка валидации" + err.Error(),
		}
	}
	createTable, err := t.tableService.CreateTable(ctx, conn, *table)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrCreateTable",
			Message: "Не удалось создать стол" + err.Error(),
		}
	}
	return createTable, nil
}

func (t *TableService) GetTableService(ctx context.Context, conn *pgx.Conn, id int) (*domain.Table, error) {
	if id <= 0 {
		return nil, errors.ValidationError{
			Field:   "id",
			Message: "ID не может быть меньше или равен нулю",
		}
	}
	table, err := t.tableService.GetByIDTable(ctx, conn, id)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}
	return table, nil
}

func (t *TableService) GetAllTablesService(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error) {
	limitOffset(limit, offset)

	tables, err := t.tableService.GetAllTable(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetAllTables",
			Message: "Не удалось вернуть все столы" + err.Error(),
		}
	}
	return tables, nil
}

func (t *TableService) GetFreeTablesService(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error) {
	limitOffset(limit, offset)

	tables, err := t.tableService.GetFreeTable(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetFreeTables",
			Message: "Не удалось вернуть все столы" + err.Error(),
		}
	}
	return tables, nil
}

func (t *TableService) GetOccupiedTablesService(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Table, error) {
	limitOffset(limit, offset)

	tables, err := t.tableService.GetOccupiedTable(ctx, conn, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetOccupiedTable",
			Message: "Не удалось вернуть все свободные столы" + err.Error(),
		}
	}
	return tables, nil
}

func (t *TableService) GetTablesByLocationService(ctx context.Context, conn *pgx.Conn, location string, limit, offset int) ([]domain.Table, error) {
	limitOffset(limit, offset)

	tables, err := t.tableService.GetByLocationTable(ctx, conn, location, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetByLocationTable",
			Message: "Не удалось вернуть все локации столов" + err.Error(),
		}
	}
	return tables, nil
}

func (t *TableService) GetTablesByCapacityService(ctx context.Context, conn *pgx.Conn, capacity, limit, offset int) ([]domain.Table, error) {
	limitOffset(limit, offset)

	tables, err := t.tableService.GetByCapacityTable(ctx, conn, capacity, limit, offset)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrGetFreeTables",
			Message: "Не удалось вернуть вместимость всех столов" + err.Error(),
		}
	}
	return tables, nil
}

func (t *TableService) UpdateTableService(ctx context.Context, conn *pgx.Conn, id int, table map[string]interface{}) (*domain.Table, error) {
	existingTable, err := t.tableService.GetByIDTable(ctx, conn, id)
	if err != nil {
		return nil, errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}

	if number, ok := table["number"].(float64); ok {
		existingTable.Number = int(number)
	}
	if capacity, ok := table["capacity"].(float64); ok {
		existingTable.Capacity = int(capacity)
	}
	if location, ok := table["location"].(string); ok {
		existingTable.Location = location
	}
	if isOccupied, ok := table["is_occupied"].(bool); ok {
		existingTable.IsOccupied = isOccupied
	}
	if isReserved, ok := table["is_reserved"].(bool); ok {
		existingTable.IsReserved = isReserved
	}

	if err := existingTable.Validate(); err != nil {
		return nil, errors.ValidationError{
			Field:   "table",
			Message: err.Error(),
		}
	}

	updated, err := t.tableService.UpdateTable(ctx, conn, id, *existingTable)
	if err != nil {
		return nil, errors.BusinessError{
			Code:    "ErrUpdateTable",
			Message: "Не удалось обновить данные стола: " + err.Error(),
		}
	}
	return updated, nil
}

func (t *TableService) DeleteTableService(ctx context.Context, conn *pgx.Conn, id int) error {
	_, err := t.tableService.GetByIDTable(ctx, conn, id)
	if err != nil {
		return errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}
	if err := t.tableService.DeleteTable(ctx, conn, id); err != nil {
		return errors.BusinessError{
			Code:    "ErrDeleteTable",
			Message: "Не удалось удалить стол " + err.Error(),
		}
	}
	return nil
}

func (t *TableService) OccupyTableService(ctx context.Context, conn *pgx.Conn, id int) error {
	table, err := t.tableService.GetByIDTable(ctx, conn, id)
	if err != nil {
		return errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}

	if table.IsOccupied {
		return errors.BusinessError{
			Code:    "ErrTableAlreadyOccupied",
			Message: "стол уже занят",
		}
	}

	if err := t.tableService.UpdateOccupiedTable(ctx, conn, id, true); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateOccupiedTable",
			Message: "Не удалось занять стол: " + err.Error(),
		}
	}

	table.IsOccupied = true
	return nil
}

func (t *TableService) FreeTableService(ctx context.Context, conn *pgx.Conn, id int) error {
	_, err := t.tableService.GetByIDTable(ctx, conn, id)
	if err != nil {
		return errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}
	if err := t.tableService.UpdateOccupiedTable(ctx, conn, id, false); err != nil {
		return errors.BusinessError{
			Code:    "ErrUpdateOccupiedTable",
			Message: "Не удалось освободить стол" + err.Error(),
		}
	}
	return nil
}

func (t *TableService) ReserveTableService(ctx context.Context, conn *pgx.Conn, id int) error {
	if _, err := t.tableService.GetByIDTable(ctx, conn, id); err != nil {
		return errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}
	if err := t.tableService.ReserveTable(ctx, conn, id); err != nil {
		return errors.BusinessError{
			Code:    "ErrReserveTable",
			Message: err.Error(),
		}
	}
	return nil
}

func (t *TableService) CancelReserveTableService(ctx context.Context, conn *pgx.Conn, id int) error {
	if _, err := t.tableService.GetByIDTable(ctx, conn, id); err != nil {
		return errors.NotFoundError{
			Entity: "table",
			ID:     id,
		}
	}
	if err := t.tableService.CancelReservation(ctx, conn, id); err != nil {
		return errors.BusinessError{
			Code:    "ErrCancelReserveTable",
			Message: "Не удалось отменить резервацию стола" + err.Error(),
		}
	}
	return nil
}
