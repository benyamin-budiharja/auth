// This file contains the repository implementation layer.
package user_repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewUserRepository(opts NewRepositoryOptions) *UserRepository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &UserRepository{
		Db: db,
	}
}
