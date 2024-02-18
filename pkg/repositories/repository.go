package repositories

import (
	"database/sql"

	"github.com/I1Asyl/task-manager-go/database"
)

type auth interface {
	GetUser(user database.UserForm) (database.User, error)
	AddSession(session database.Session) error
	CheckRefreshToken(first_token, token string) (bool, error)
	DeleteToken(first_token string) error
	GetUserByFirstToken(first_token string) (database.User, error)
	UpdateToken(first_token string, token string) error
}

type admin interface {
	CreateUser(user database.User) error
	CreateTeam(team database.Team) error
	AddUserToTeam(userId int, teamId int, roleId int) error
	GetTeamMembers(teamId int) ([]database.User, error)
	CanEditTeamUser(userId int, teamId int) (bool, error)
}

type static interface {
	ReturnStatus() string
}

// Repository structure to c
type Repository struct {
	auth
	static
	admin
}

// New returns a new repository with relevant methods configured
func New(db *sql.DB) *Repository {
	return &Repository{
		auth:   NewAuthorization(db),
		static: NewStatic(db),
		admin:  NewAdmin(db),
	}
}

// use sqlc for generating db access code in Go
