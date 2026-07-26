package singleton

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConnection struct {
	pool *pgxpool.Pool
}

var (
	dbInstance *DBConnection
	dbOnce     sync.Once
)

func GetDBConnection(dsn string) *DBConnection {
	dbOnce.Do(func() {
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			panic("failed to parse DSN: " + err.Error())
		}

		config.MaxConns = 25
		config.MinConns = 5
		config.MaxConnLifetime = 1 * time.Hour
		config.MaxConnIdleTime = 30 * time.Minute

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			panic("failed to create connection pool: " + err.Error())
		}

		if err := pool.Ping(context.Background()); err != nil {
			panic("failed to ping database: " + err.Error())
		}

		dbInstance = &DBConnection{pool: pool}
	})

	return dbInstance
}

func (db *DBConnection) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}

func (db *DBConnection) Ping(ctx context.Context) error {
	if db.pool == nil {
		return fmt.Errorf("connection pool is not initialized")
	}
	return db.pool.Ping(ctx)
}

// возвращает пул подключений
func (db *DBConnection) GetPool() *pgxpool.Pool {
	return db.pool
}
