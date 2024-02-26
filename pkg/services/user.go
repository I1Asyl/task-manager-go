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

func (a User) CreateProject(model database.Model) (map[string]string, error) {
	project := database.Project(model.Project)
	if mistakes := project.IsValid(1); len(mistakes) > 0 {
		return mistakes, nil
	}
	ok, err := a.repo.CanEditTeamProject(model.CurrentUser.Id, model.Project.TeamId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return map[string]string{"permission": "user can't edit project"}, nil
	}

	return nil, a.repo.CreateProject(project)
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

func (a User) CreateTask(model database.Model) (map[string]string, error) {
	task := database.Task(model.Task)
	task.AssignerId = model.CurrentUser.Id
	if mistakes := task.IsValid(1); len(mistakes) > 0 {
		return mistakes, nil
	}
	teamId, err := a.repo.GetTeamByProjectId(task.ProjectId)
	if err != nil {
		return nil, err
	}
	model.Team.Id = teamId
	ok, err := a.repo.CanEditTeamProject(model.CurrentUser.Id, model.Team.Id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return map[string]string{"permission": "user can't edit project"}, nil
	}
	ok, err = a.repo.IsInTeam(task.UserId, model.Team.Id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return map[string]string{"user": "added user is not in the team"}, nil
	}

	return nil, a.repo.CreateTask(task)
}

func (a User) GetTasksByProject(model database.Model) ([]database.Task, error) {
	project := database.Project(model.Project)
	project.TeamId, _ = a.repo.GetTeamByProjectId(project.Id)

	ok, err := a.repo.IsInTeam(model.CurrentUser.Id, project.TeamId)
	if err != nil {
		return []database.Task{}, err
	}
	if !ok {
		return []database.Task{}, errors.New("user can't see tasks")
	}
	return a.repo.GetTasksByProject(project.Id)
}

func (a User) GetTasks(model database.Model) ([]database.Task, error) {
	return a.repo.GetTasks(model.CurrentUser.Id)
}

func (a User) UpdateTask(model database.Model) (map[string]string, error) {
	task := database.Task(model.Task)
	if mistakes := task.IsValid(2); len(mistakes) > 0 {
		return mistakes, nil
	}

	ok, err := a.repo.CanEditTask(model.CurrentUser.Id, task.Id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("user can't edit task")
	}
	return nil, a.repo.UpdateTask(task)
}
