package repository

import (
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(userEmail, password string) (todo.User, error)
}

type RequestList interface {
}

type Repository struct {
	Authorization
	RequestList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
