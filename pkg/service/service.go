package service

import (
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"github.com/TelitsynNikita/bookkeeper-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type RequestList interface {
}

type Service struct {
	Authorization
	RequestList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
