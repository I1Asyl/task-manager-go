package services

import (
	"fmt"

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
	fmt.Println(user.Name)
	if mistakes := user.IsValid(1); len(mistakes) > 0 {
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
	if mistakes := team.IsValid(1); len(mistakes) > 0 {
		fmt.Println(mistakes)
	}

	if err := a.repo.CreateTeam(team); err != nil {
		return nil, err
	}
	return nil, nil
}

func (a Admin) DeleteTeam(model database.Model) error {
	return a.repo.DeleteTeam(model.Team.Id)
}
