package services

import "github.com/I1Asyl/task-manager-go/pkg/repositories"

type Static struct {
	repo *repositories.Repository
}

func NewStatic(repo *repositories.Repository) *Static {
	return &Static{repo: repo}
}

func (s *Static) ReturnStatus() string {
	return s.repo.ReturnStatus()
}
