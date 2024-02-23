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
	userTeam := database.UserTeam(model.UserTeam)
	ok, err := a.repo.CanEditTeamUser(model.CurrentUser.Id, userTeam.TeamId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user can't edit team")
	}

	return a.repo.AddUserToTeam(userTeam.UserId, userTeam.TeamId, userTeam.RoleId)
}

func (a User) CreateProject(model database.Model) error {
	project := database.Project(model.Project)
	if mistakes := project.IsValid(); len(mistakes) > 0 {
		for _, m := range mistakes {
			return errors.New(m)
		}
	}
	ok, err := a.repo.CanEditTeamProject(model.CurrentUser.Id, model.Project.TeamId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user can't edit project")
	}

	return a.repo.CreateProject(project)
}

func (a User) GetTeamMembers(model database.Model) ([]database.User, error) {
	team := database.Team(model.Team)

	// ok, err := a.repo.CanEditTeamUser(model.CurrentUser.Id, team.Id)
	// fmt.Println(model.CurrentUser.Id, team.Id, ok, err)
	// if err != nil {
	// 	return []database.User{}, err
	// }
	// if !ok {
	// 	return []database.User{}, errors.New("user can't edit or view team")
	// }
	return a.repo.GetTeamMembers(team.Id)
}

func (a User) CreateTask(model database.Model) error {
	task := database.Task(model.Task)
	task.AssignerId = model.CurrentUser.Id
	if mistakes := task.IsValid(); len(mistakes) > 0 {
		for _, m := range mistakes {
			return errors.New(m)
		}
	}
	teamId, err := a.repo.GetTeamByProjectId(task.ProjectId)
	if err != nil {
		return err
	}
	model.Team.Id = teamId
	ok, err := a.repo.CanEditTeamProject(model.CurrentUser.Id, model.Team.Id)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("user can't edit project")
	}

	return a.repo.CreateTask(task)
}
