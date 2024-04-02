// Package db contain CRUD operations to be called from handlers
package db

import "github.com/jackc/pgx/v5"

type PostgresRepo struct {
	Client *pgx.Conn
}
