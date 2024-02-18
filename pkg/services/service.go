package services

import (
	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
)

type auth interface {
	Login(model database.Model) (string, string, map[string]string, error)
	VerifyUser(tokenString string) bool
	VerifyAdmin(tokenString string) bool
	RefreshToken(tokenString string) (string, string, error)
	GetUserIdByToken(token string) (int, error)
	Logout(token string) error
}

type admin interface {
	CreateUser(model database.Model) (map[string]string, error)
	CreateTeam(model database.Model) (map[string]string, error)
}

type user interface {
	AddUserToTeam(model database.Model) error
	GetTeamMembers(model database.Model) ([]database.User, error)
}

type Service struct {
	auth
	user
	admin
}

func New(repo *repositories.Repository) *Service {
	return &Service{auth: NewAuthorization(repo), user: NewUser(repo), admin: NewAdmin(repo)}
}
