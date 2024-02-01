package services

import "github.com/I1Asyl/task-manager-go/pkg/repositories"

type Service struct {
	repo *repositories.Repository
}

func New(repo *repositories.Repository) Service {
	return Service{repo: repo}
}

func (s *Service) ReturnStatus() string {
	return s.repo.ReturnStatus()
}
