package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/repository/gen"
)

type Store struct {
	queries *gen.Queries
	conn    *pgxpool.Pool
}

func NewStore(conn *pgxpool.Pool) *Store {

	queries := gen.New(conn)
	return &Store{
		queries: queries,
		conn:    conn,
	}
}

func OpenDB(url string) (*pgxpool.Pool, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, func() {}, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, func() {}, err
	}

	return conn, func() {
		conn.Close()
	}, nil
}
