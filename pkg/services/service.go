package services

import (
	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
)

type auth interface {
	CreateUser(model database.Model) (map[string]string, error)
	Login(model database.Model) (string, string, map[string]string, error)
	VerifyUser(tokenString string) bool
	VerifyAdmin(tokenString string) bool
	RefreshToken(tokenString string) (string, string, error)
	CreateTeam(model database.Model) (map[string]string, error)
	GetUserIdByToken(token string) (int, error)
	Logout(token string) error
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
