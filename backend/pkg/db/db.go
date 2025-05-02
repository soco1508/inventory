package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func SqlxConnect(ctx context.Context, dbConfig DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", getConnectionStr(dbConfig))
	if err != nil {
		return nil, fmt.Errorf("cannot connect sqlxdb, err:\n %+v", err)
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("cannot ping sqlxdb, err:\n %+v", err)
	}
	return db, nil
}

func PgxConnect(ctx context.Context, dbConfig DBConfig) (*pgx.Conn, error) {
	db, err := pgx.Connect(ctx, getConnectionStr(dbConfig))
	if err != nil {
		return nil, fmt.Errorf("cannot connect pgxdb, err:\n %+v", err)
	}
	return db, nil
}

func getConnectionStr(dbConfig DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name,
	)
}
