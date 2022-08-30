package service

import (
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"github.com/TelitsynNikita/bookkeeper-backend/pkg/repository"
)

type RequestListService struct {
	repo repository.RequestList
}

func NewRequestListService(repo repository.RequestList) *RequestListService {
	return &RequestListService{repo: repo}
}

func (s *RequestListService) Create(userId int, request todo.Request) (int, error) {
	return s.repo.Create(userId, request)
}

func (s *RequestListService) GetAll(userId int) ([]todo.AllRequests, error) {
	return s.repo.GetAll(userId)
}

func (s *RequestListService) GetOne(requestId int) (todo.OneRequest, error) {
	return s.repo.GetOne(requestId)
}
