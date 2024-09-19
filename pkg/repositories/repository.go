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
	DeleteTeam(teamId int) error
}

type user interface {
	GetTeamMembers(teamId int) ([]database.User, error)
	AddUserToTeam(userId int, teamId int, roleId int) error
	CanEditTeamUser(userId int, teamId int) (bool, error)
	CanEditTeamProject(userId int, teamId int) (bool, error)
	CreateProject(project database.Project) error
	CreateTask(task database.Task) error
	GetTeamByProjectId(projectId int) (int, error)
	IsInTeam(userId int, teamId int) (bool, error)
	GetTasksByProject(projectId int) ([]database.Task, error)
	GetTasks(userId int) ([]database.Task, error)
	UpdateTask(task database.Task) error
	Update(tablename string, allColumnNames []string, allColumnValues []interface{}, id int) error
	CanEditTask(userId int, taskId int) (bool, error)
	UpdateProject(project database.Project) error
	GetProjects(projectId int) ([]database.Project, error)
	GetProject(projectId int) (database.Project, error)
	GetTask(taskId int) (database.Task, error)
}

type Repository struct {
	auth
	admin
	user
}

// New returns a new repository with relevant methods configured
func New(db *sql.DB) *Repository {
	return &Repository{
		auth:  NewAuthorization(db),
		admin: NewAdmin(db),
		user:  NewUser(db),
	}
}

// use sqlc for generating db access code in Go
