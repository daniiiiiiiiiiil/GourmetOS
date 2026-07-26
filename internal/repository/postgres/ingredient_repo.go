package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IngredientRepository struct {
	Pool *pgxpool.Pool
}

func NewIngredientRepository(pool *pgxpool.Pool) *IngredientRepository {
	return &IngredientRepository{Pool: pool}
}

func (i *IngredientRepository) CreateIngredient(ctx context.Context, conn *pgx.Conn, ingredient domain.Ingredient) (*domain.Ingredient, error) {
	sqlQuery := `
	INSERT INTO ingredients(name, unit, stock_quantity, min_stock, max_stock, cost_per_unit, supplier, expiry_date, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING ingredient_id, created_at, updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		ingredient.Name,
		ingredient.Unit,
		ingredient.StockQuantity,
		ingredient.MinStock,
		ingredient.MaxStock,
		ingredient.CostPerUnit,
		ingredient.Supplier,
		ingredient.ExpiryDate,
		time.Now(),
		time.Now(),
	).Scan(&ingredient.IngredientID, &ingredient.CreatedAt, &ingredient.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (i *IngredientRepository) GetByIDIngredient(ctx context.Context, conn *pgx.Conn, ingredientID int) (*domain.Ingredient, error) {
	sqlQuery := `
	SELECT *
	FROM ingredients
	WHERE ingredient_id = $1`

	rows, err := conn.Query(ctx, sqlQuery, ingredientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredient, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return &ingredient, nil
}

func (i *IngredientRepository) GetAllIngredients(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Ingredient, error) {
	sqlQuery := `
	SELECT *
	FROM ingredients
	LIMIT $1 OFFSET $2`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (i *IngredientRepository) GetByNameIngredient(ctx context.Context, conn *pgx.Conn, name string, limit, offset int) ([]domain.Ingredient, error) {
	sqlQuery := `
	SELECT *
	FROM ingredients
	WHERE name ILIKE $1
	LIMIT $2 OFFSET $3`

	rows, err := conn.Query(ctx, sqlQuery, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (i *IngredientRepository) UpdateIngredient(ctx context.Context, conn *pgx.Conn, ingredient domain.Ingredient, id int) (*domain.Ingredient, error) {
	sqlQuery := `
	UPDATE ingredients
	SET name = $1, unit = $2, stock_quantity = $3, min_stock = $4,
		max_stock = $5, cost_per_unit = $6, supplier = $7, expiry_date = $8,
		updated_at = $9
	WHERE ingredient_id = $10
	RETURNING ingredient_id, name, unit, stock_quantity, min_stock, max_stock,
			  cost_per_unit, supplier, expiry_date, created_at, updated_at`

	rows, err := conn.Query(ctx, sqlQuery,
		ingredient.Name,
		ingredient.Unit,
		ingredient.StockQuantity,
		ingredient.MinStock,
		ingredient.MaxStock,
		ingredient.CostPerUnit,
		ingredient.Supplier,
		ingredient.ExpiryDate,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	updated, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (i *IngredientRepository) UpdateStockIngredient(ctx context.Context, conn *pgx.Conn, id int, stockQuantity float64) error {
	sqlQuery := `
	UPDATE ingredients
	SET stock_quantity = $1, updated_at = $2
	WHERE ingredient_id = $3`

	_, err := conn.Exec(ctx, sqlQuery, stockQuantity, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (i *IngredientRepository) GetLowStockIngredients(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Ingredient, error) {
	sqlQuery := `
	SELECT *
	FROM ingredients
	WHERE stock_quantity <= min_stock
	LIMIT $1 OFFSET $2`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (i *IngredientRepository) DeleteIngredient(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `DELETE FROM ingredients WHERE ingredient_id = $1`

	_, err := conn.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (i *IngredientRepository) GetBySupplierIngredient(ctx context.Context, conn *pgx.Conn, supplier string, limit, offset int) ([]domain.Ingredient, error) {
	sqlQuery := `
	SELECT *
	FROM ingredients
	WHERE supplier ILIKE $1
	LIMIT $2 OFFSET $3`

	rows, err := conn.Query(ctx, sqlQuery, "%"+supplier+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (i *IngredientRepository) GetExpiringSoonIngredients(ctx context.Context, conn *pgx.Conn, days int, limit, offset int) ([]domain.Ingredient, error) {
	sqlQuery := `
	SELECT *
	FROM ingredients
	WHERE expiry_date <= CURRENT_DATE + $1
	ORDER BY expiry_date ASC
	LIMIT $2 OFFSET $3`

	rows, err := conn.Query(ctx, sqlQuery, days, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Ingredient])
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}
