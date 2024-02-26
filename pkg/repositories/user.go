package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/I1Asyl/task-manager-go/database"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db}
}

func (a User) AddUserToTeam(userId int, teamId int, roleId int) error {
	_, err := a.db.Query("INSERT INTO users_teams (user_id, team_id, role_id) VALUES ($1, $2, $3) RETURNING *", userId, teamId, roleId)
	return err
}

func (a User) GetTeamMembers(teamId int) ([]database.User, error) {
	users := []database.User{}
	res, err := a.db.Query("SELECT id, username, name, surname, phone, email, is_admin FROM users WHERE id IN (SELECT user_id FROM users_teams WHERE team_id = $1);", teamId)
	for res.Next() {
		var user database.User
		res.Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Phone, &user.Email, &user.IsAdmin)
		users = append(users, user)
	}

	return users, err
}

func (a User) CanEditTeamUser(userId int, teamId int) (bool, error) {
	var canEdit bool
	err := a.db.QueryRow("SELECT can_edit_users FROM team_roles WHERE id IN (SELECT role_id FROM users_teams WHERE user_id = $1 AND team_id = $2)", userId, teamId).Scan(&canEdit)
	return canEdit, err
}

func (a User) CanEditTeamProject(userId int, teamId int) (bool, error) {
	var canEdit bool
	err := a.db.QueryRow("SELECT can_edit_projects FROM team_roles WHERE id IN (SELECT role_id FROM users_teams WHERE user_id = $1 AND team_id = $2)", userId, teamId).Scan(&canEdit)
	return canEdit, err
}

func (a User) CreateProject(project database.Project) error {
	_, err := a.db.Query("INSERT INTO projects (name, description, team_id, current_status) VALUES ($1, $2, $3, $4)", project.Name, project.Description, project.TeamId, project.CurrentStatus)
	return err
}

func (a User) IsInTeam(userId int, teamId int) (bool, error) {
	var isInTeam bool
	err := a.db.QueryRow("SELECT EXISTS (SELECT 1 FROM users_teams WHERE user_id = $1 AND team_id = $2)", userId, teamId).Scan(&isInTeam)
	return isInTeam, err
}

func (a User) CreateTask(task database.Task) error {
	_, err := a.db.Query("INSERT INTO tasks (name, description, project_id, current_status, user_id, assigner_id, start_time) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", task.Name, task.Description, task.ProjectId, task.CurrentStatus, task.UserId, task.AssignerId, time.Now())

	return err
}

func (a User) GetTeamByProjectId(projectId int) (int, error) {
	var teamId int
	err := a.db.QueryRow("SELECT team_id FROM projects WHERE id = $1", projectId).Scan(&teamId)
	return teamId, err
}

func (a User) GetTasksByProject(projectId int) ([]database.Task, error) {
	var tasks []database.Task
	res, err := a.db.Query("SELECT id, name, description, project_id, current_status, assigner_id, start_time FROM tasks WHERE project_id = $1", projectId)
	for res.Next() {
		var task database.Task
		res.Scan(&task.Id, &task.Name, &task.Description, &task.ProjectId, &task.CurrentStatus, &task.AssignerId, &task.StartTime)
		tasks = append(tasks, task)
	}

	return tasks, err
}

func (a User) GetTasksByTeam(teamId int) ([]database.Task, error) {
	var tasks []database.Task
	res, err := a.db.Query("SELECT tasks.id, tasks.name, tasks.description, tasks.project_id, tasks.current_status, tasks.assigner_id, tasks.start_time FROM tasks JOIN projects ON tasks.project_id = projects.id WHERE projects.team_id = $1", teamId)
	for res.Next() {
		var task database.Task
		res.Scan(&task.Id, &task.Name, &task.Description, &task.ProjectId, &task.CurrentStatus, &task.AssignerId, &task.StartTime)
		tasks = append(tasks, task)
	}

	return tasks, err
}

func (a User) GetTasks(userId int) ([]database.Task, error) {
	var tasks []database.Task
	res, err := a.db.Query("SELECT id, name, description, project_id, current_status, assigner_id, start_time FROM tasks WHERE user_id = $1", userId)
	for res.Next() {
		var task database.Task
		res.Scan(&task.Id, &task.Name, &task.Description, &task.ProjectId, &task.CurrentStatus, &task.AssignerId, &task.StartTime)
		tasks = append(tasks, task)
	}

	return tasks, err
}

func (a User) Update(tablename string, allColumnNames []string, allColumnValues []interface{}, id int) error {
	// allColumnNames := []string{"name", "description", "current_status", "assigner_id", "start_time"}
	// allColumnValues := []interface{}{task.Name, task.Description, task.CurrentStatus, task.AssignerId, task.StartTime}
	columnNames := []string{}
	columnValues := []interface{}{}
	for i, columnValue := range allColumnValues {
		if columnValue != 0 && columnValue != "" {
			columnNames = append(columnNames, allColumnNames[i])
			columnValues = append(columnValues, columnValue)
		}
	}

	query := fmt.Sprintf("UPDATE %s SET ", tablename)
	for i, columnName := range columnNames {
		query += columnName + " = $" + fmt.Sprint(i+1)
		if i < len(columnNames)-1 {
			query += ", "
		}
	}
	query += " WHERE id = $" + fmt.Sprint(len(columnValues)+1)

	columnValues = append(columnValues, id)
	_, err := a.db.Query(query, columnValues...)
	return err
}

func (a User) UpdateTask(task database.Task) error {
	allColumnNames := []string{"name", "description", "current_status", "assigner_id", "start_time"}
	allColumnValues := []interface{}{task.Name, task.Description, task.CurrentStatus, task.AssignerId, task.StartTime}
	err := a.Update("tasks", allColumnNames, allColumnValues, task.Id)
	return err
}

func (a User) CanEditTask(userId int, taskId int) (bool, error) {
	var canEdit bool
	err := a.db.QueryRow("SELECT EXISTS (SELECT 1 FROM tasks WHERE id = $2 AND (assigner_id = $1 OR user_id = $1))", userId, taskId).Scan(&canEdit)
	return canEdit, err
}

func (a User) UpdateProject(project database.Project) error {
	allColumnNames := []string{"name", "description", "current_status"}
	allColumnValues := []interface{}{project.Name, project.Description, project.CurrentStatus}
	err := a.Update("projects", allColumnNames, allColumnValues, project.Id)
	return err
}

func (a User) GetProjects(userId int) ([]database.Project, error) {
	var projects []database.Project
	res, err := a.db.Query("SELECT id, name, description, team_id, current_status FROM projects WHERE team_id in (SELECT team_id FROM users_teams WHERE user_id = $1)", userId)
	for res.Next() {
		var project database.Project
		res.Scan(&project.Id, &project.Name, &project.Description, &project.TeamId, &project.CurrentStatus)
		projects = append(projects, project)
	}
	return projects, err
}

func (a User) GetProject(projectId int) (database.Project, error) {
	var project database.Project
	err := a.db.QueryRow("SELECT id, name, description, team_id, current_status FROM projects WHERE id = $1", projectId).Scan(&project.Id, &project.Name, &project.Description, &project.TeamId, &project.CurrentStatus)
	return project, err
}

func (a User) GetTask(taskId int) (database.Task, error) {
	var task database.Task
	err := a.db.QueryRow("SELECT id, name, description, project_id, current_status, assigner_id, start_time FROM tasks WHERE id = $1", taskId).Scan(&task.Id, &task.Name, &task.Description, &task.ProjectId, &task.CurrentStatus, &task.AssignerId, &task.StartTime)
	return task, err
}
