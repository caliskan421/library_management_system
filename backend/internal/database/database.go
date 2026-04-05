package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/muhammetalicaliskan/libranet/internal/config"
	"github.com/muhammetalicaliskan/libranet/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New(cfg config.DBConfig) (*bun.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db, nil
}

func Migrate(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*model.User)(nil),
		(*model.Book)(nil),
		(*model.Reservation)(nil),
	}

	for _, m := range models {
		if _, err := db.NewCreateTable().
			Model(m).
			IfNotExists().
			Exec(ctx); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	return nil
}
