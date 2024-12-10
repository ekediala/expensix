package sqlx

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func Dial(ctx context.Context, address string) (*DB, error) {
	db, err := sql.Open("postgres", address)
	if err != nil {
		return nil, fmt.Errorf("opening database %s: %w", address, err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging db %s:%w", address, err)
	}

	return &DB{DB: db}, nil
}
