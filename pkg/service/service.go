package service

import (
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"github.com/TelitsynNikita/bookkeeper-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(userEmail, password string) (string, error)
	ParseToken(token string) (int, error)
}

type RequestList interface {
	Create(userId int, request todo.Request) (int, error)
	GetAll(userId int) ([]todo.AllRequests, error)
	GetOne(requestId int) (todo.OneRequest, error)
}

type Service struct {
	Authorization
	RequestList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		RequestList:   NewRequestListService(repos.RequestList),
	}
}
