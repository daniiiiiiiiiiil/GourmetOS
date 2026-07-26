package postgres

import (
	"GourmetOS/internal/domain"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DishRepository struct {
	Pool *pgxpool.Pool
}

func NewDishRepository(pool *pgxpool.Pool) *DishRepository {
	return &DishRepository{Pool: pool}
}

func (d *DishRepository) CreateDish(ctx context.Context, conn *pgx.Conn, dish domain.Dish) (*domain.Dish, error) {
	sqlQuery := `
	INSERT INTO dish (name, description, price, category, cuisine, cooking_time,
					  is_available, is_vegetarian, is_vegan, is_gluten_free,
					  calories, image_url, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING dish_id, created_at, updated_at`

	err := conn.QueryRow(ctx, sqlQuery,
		dish.Name,
		dish.Description,
		dish.Price,
		dish.Category,
		dish.Cuisine,
		dish.CookingTime,
		dish.IsAvailable,
		dish.IsVegetarian,
		dish.IsVegan,
		dish.IsGlutenFree,
		dish.Calories,
		dish.ImageURL,
		time.Now(),
	).Scan(&dish.DishID, &dish.CreatedAt, &dish.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &dish, nil
}

func (d *DishRepository) GetByIDDish(ctx context.Context, conn *pgx.Conn, dishID int) (*domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE dish_id = $1`

	rows, err := conn.Query(ctx, sqlQuery, dishID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dish, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}

	return &dish, nil
}

func (d *DishRepository) GetAllDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
		SELECT * FROM dish
		LIMIT $1 OFFSET $2
		`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	return dishes, nil
}

func (d *DishRepository) UpdateDish(ctx context.Context, conn *pgx.Conn, id int, dish domain.Dish) (*domain.Dish, error) {
	sqlQuery := `
	UPDATE dish
	SET name = $1, description = $2, price = $3, category = $4, cuisine = $5,
		cooking_time = $6,is_available = $7, is_vegetarian = $8, is_vegan = $9,
		is_gluten_free = $10,calories = $11,image_url = $12,updated_at = $13
		WHERE dish_id = $14`
	_, err := conn.Exec(ctx, sqlQuery,
		dish.Name,
		dish.Description,
		dish.Price,
		dish.Category,
		dish.Cuisine,
		dish.CookingTime,
		dish.IsAvailable,
		dish.IsVegetarian,
		dish.IsVegan,
		dish.IsGlutenFree,
		dish.Calories,
		dish.ImageURL,
		time.Now(),
		id)
	if err != nil {
		return nil, err
	}
	dish.DishID = id
	return &dish, nil
}

func (d *DishRepository) DeleteDish(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `DELETE FROM dish WHERE dish_id = $1`
	_, err := conn.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepository) GetByCategoryDish(ctx context.Context, conn *pgx.Conn, category string, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE category = $1
	LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, category, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishRepository) GetByCuisineDish(ctx context.Context, conn *pgx.Conn, cuisine string, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE cuisine = $1
	LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, cuisine, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishRepository) GetByNameDish(ctx context.Context, conn *pgx.Conn, name string, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE name = $1
	LIMIT $2 OFFSET $3`
	rows, err := conn.Query(ctx, sqlQuery, name, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishRepository) GetByAvailable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE is_available = true
	LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishRepository) UpdateAvailable(ctx context.Context, conn *pgx.Conn, id int, isAvailable bool) error {
	sqlQuery := `
	UPDATE dish
	SET is_available = $2
	WHERE dish_id = $1`
	_, err := conn.Exec(ctx, sqlQuery, id, isAvailable)
	if err != nil {
		return err
	}
	return nil
}

func (d *DishRepository) GetByPriceRange(ctx context.Context, conn *pgx.Conn, min, max, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE price BETWEEN $1 AND $2
	LIMIT $3 OFFSET $4`
	rows, err := conn.Query(ctx, sqlQuery, min, max, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishRepository) GetVegetarianDish(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE is_vegetarian = true
	LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	return dishes, nil
}

func (d *DishRepository) GetVeganDish(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE is_vegan = true
	LIMIT $1 OFFSET $2`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (d *DishRepository) GetGlutenFreeDish(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error) {
	sqlQuery := `
	SELECT * FROM dish
	WHERE is_gluten_free = true
	LIMIT $1 OFFSET $2`

	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Dish])
	if err != nil {
		return nil, err

	}
	return dishes, nil
}
