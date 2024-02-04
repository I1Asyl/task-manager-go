package services

import "github.com/I1Asyl/task-manager-go/pkg/repositories"

type auth interface {
}

type static interface {
	ReturnStatus() string
}

type Service struct {
	auth
	static
}

func New(repo *repositories.Repository) *Service {
	return &Service{auth: NewAuthorization(repo), static: NewStatic(repo)}
}
