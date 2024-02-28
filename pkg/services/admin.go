package services

import (
	"crypto/sha256"
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

func Hash(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	ans := h.Sum([]byte("secret"))
	res := fmt.Sprintf("%x", ans)
	return res
}

func (a Admin) CreateUser(model database.Model) (map[string]string, error) {

	user := database.User(model.User)
	if mistakes := user.IsValid(1); len(mistakes) > 0 {
		return mistakes, nil
	}
	user.Password = Hash(user.Password)
	err := a.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a Admin) CreateTeam(model database.Model) (map[string]string, error) {

	team := database.Team(model.Team)
	if mistakes := team.IsValid(1); len(mistakes) > 0 {
		return mistakes, nil
	}

	if err := a.repo.CreateTeam(team); err != nil {
		return nil, err
	}
	return nil, nil
}

func (a Admin) DeleteTeam(model database.Model) error {
	return a.repo.DeleteTeam(model.Team.Id)
}
