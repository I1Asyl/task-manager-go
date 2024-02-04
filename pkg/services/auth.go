package services

import "github.com/I1Asyl/task-manager-go/pkg/repositories"

type Authorization struct {
	repo *repositories.Repository
}

func NewAuthorization(repo *repositories.Repository) *Authorization {
	return &Authorization{repo: repo}
}
