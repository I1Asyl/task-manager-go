package repositories

import (
	"database/sql"

	"github.com/I1Asyl/task-manager-go/database"
)

type auth interface {
	CreateUser(user database.User) error
	GetUser(user database.UserForm) (database.User, error)
	AddSession(session database.Session) error
	CheckRefreshToken(first_token, token string) (bool, error)
	DeleteToken(first_token string) error
	GetUserByFirstToken(first_token string) (database.User, error)
	UpdateToken(first_token string, token string) error
	CreateTeam(team database.Team) error
}

type static interface {
	ReturnStatus() string
}

// Repository structure to c
type Repository struct {
	auth
	static
}

// New returns a new repository with relevant methods configured
func New(db *sql.DB) *Repository {
	return &Repository{
		auth:   NewAuthorization(db),
		static: NewStatic(db),
	}
}

// use sqlc for generating db access code in Go
