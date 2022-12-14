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

func (s *RequestListService) GetAll() ([]todo.AllRequests, error) {
	return s.repo.GetAll()
}

func (s *RequestListService) GetOne(requestId int) (todo.OneRequest, error) {
	return s.repo.GetOne(requestId)
}

func (s *RequestListService) DeleteOne(requestId int) error {
	return s.repo.DeleteOne(requestId)
}

func (s *RequestListService) Update(requestStatus todo.UpdateRequestStatus) error {
	return s.repo.Update(requestStatus)
}

func (s *RequestListService) UpdateRequestBookkeeperId(requestBookkeeperId todo.UpdateRequestBookkeeperId) error {
	return s.repo.UpdateBookkeeperId(requestBookkeeperId)
}
