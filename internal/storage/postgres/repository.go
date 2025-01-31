package db

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/ShayeGun/go-server/internal/common"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	userCollection common.UserRepositoryInterface
	q              *Queries
}

func NewRepository(ctx context.Context, connStr string) (*Repository, error) {

	conn, err := pgx.Connect(ctx, os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error in connection")
	}

	q := New(conn)

	return &Repository{
		q:              q,
		userCollection: NewUserRepository(ctx, q),
	}, nil
}

func (r *Repository) GetUserTable() common.UserRepositoryInterface {
	return r.userCollection
}
