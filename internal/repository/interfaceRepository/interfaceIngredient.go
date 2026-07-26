package interfaceRepository

import (
	"GourmetOS/internal/domain"
	"context"

	"github.com/jackc/pgx/v5"
)

type IngredientRepository interface {
	CreateIngredient(ctx context.Context, conn *pgx.Conn, ingredient domain.Ingredient) (*domain.Ingredient, error)
	GetByIDIngredient(ctx context.Context, conn *pgx.Conn, ingredientID int) (*domain.Ingredient, error)
	GetAllIngredients(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Ingredient, error)
	GetByNameIngredient(ctx context.Context, conn *pgx.Conn, name string, limit, offset int) ([]domain.Ingredient, error)
	UpdateIngredient(ctx context.Context, conn *pgx.Conn, ingredient domain.Ingredient, id int) (*domain.Ingredient, error)
	UpdateStockIngredient(ctx context.Context, conn *pgx.Conn, id int, stockQuantity float64) error
	GetLowStockIngredients(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]domain.Ingredient, error)
	DeleteIngredient(ctx context.Context, conn *pgx.Conn, id int) error
	GetBySupplierIngredient(ctx context.Context, conn *pgx.Conn, supplier string, limit, offset int) ([]domain.Ingredient, error)
	GetExpiringSoonIngredients(ctx context.Context, conn *pgx.Conn, days int, limit, offset int) ([]domain.Ingredient, error)
}
