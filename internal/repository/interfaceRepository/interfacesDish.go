package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type DishRepository interface {
	CreateDish(ctx context.Context, conn *pgx.Conn, dish domain.Dish) (*domain.Dish, error)
	GetByIDDish(ctx context.Context, conn *pgx.Conn, dishID int) (*domain.Dish, error)
	GetAllDishes(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error)
	UpdateDish(ctx context.Context, conn *pgx.Conn, id int, dish domain.Dish) (*domain.Dish, error)
	DeleteDish(ctx context.Context, conn *pgx.Conn, id int) error
	GetByCategoryDish(ctx context.Context, conn *pgx.Conn, category string, limit, offset int) ([]domain.Dish, error)
	GetByCuisineDish(ctx context.Context, conn *pgx.Conn, cuisine string, limit, offset int) ([]domain.Dish, error)
	GetByNameDish(ctx context.Context, conn *pgx.Conn, name string, limit, offset int) ([]domain.Dish, error)
	GetByAvailable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error)
	UpdateAvailable(ctx context.Context, conn *pgx.Conn, dishID int, isAvailable bool) error
	GetByPriceRange(ctx context.Context, conn *pgx.Conn, min, max, limit, offset int) ([]domain.Dish, error)
	GetVegetarianDish(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error)
	GetVeganDish(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error)
	GetGlutenFreeDish(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Dish, error)
}
