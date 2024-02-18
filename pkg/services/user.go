package services

import (
	"errors"

	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
)

type User struct {
	repo *repositories.Repository
}

func NewUser(repo *repositories.Repository) *User {
	return &User{repo: repo}
}

func (a User) AddUserToTeam(model database.Model) error {
	team := database.Team(model.Team)
	user := database.User(model.User)
	role := database.Role(model.Role)
	ok, err := a.repo.CanEditTeamUser(model.CurrentUser.Id, team.Id)
	if !ok {
		return errors.New("user can't edit team")
	}
	if err != nil {
		return err
	}
	return a.repo.AddUserToTeam(user.Id, team.Id, role.Id)
}

func (a User) CreateProject(model database.Model) error {
	project := database.Project(model.Project)
	ok, err := a.repo.CanEditTeamProject(model.CurrentUser.Id, model.Team.Id)
	if !ok {
		return errors.New("user can't edit project")
	}
	if err != nil {
		return err
	}
	return a.repo.CreateProject(project, model.Team.Id)
}

func (a User) GetTeamMembers(model database.Model) ([]database.User, error) {
	team := database.Team(model.Team)
	return a.repo.GetTeamMembers(team.Id)
}
