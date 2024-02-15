package services

import (
	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
)

type Admin struct {
	repo *repositories.Repository
}

func NewAdmin(repo *repositories.Repository) *Admin {
	return &Admin{repo: repo}
}

func (a Admin) CreateUser(model database.Model) (map[string]string, error) {

	user := database.User(model.User)
	if mistakes := user.IsValid(); len(mistakes) > 0 {
		return mistakes, nil
	}
	err := a.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a Admin) CreateTeam(model database.Model) (map[string]string, error) {

	team := database.Team(model.Team)
	if mistakes := team.IsValid(); len(mistakes) > 0 {
		return mistakes, nil
	}
	err := a.repo.CreateTeam(team)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a Admin) AddUserToTeam(model database.Model) error {
	team := database.Team(model.Team)
	user := database.User(model.User)
	return a.repo.AddUserToTeam(user, team)
}
