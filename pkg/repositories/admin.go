package repositories

import (
	"database/sql"

	"github.com/I1Asyl/task-manager-go/database"
)

type Admin struct {
	db *sql.DB
}

func NewAdmin(db *sql.DB) *Admin {
	return &Admin{
		db: db,
	}
}

func (a Admin) CreateUser(user database.User) error {
	_, err := a.db.Query("INSERT INTO users (username, password, name, surname, phone, email, is_admin) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.Username, user.Password, user.Name, user.Surname, user.Phone, user.Email, user.IsAdmin)
	return err
}

func (a Admin) CreateTeam(team database.Team) error {
	_, err := a.db.Query("INSERT INTO teams (name) VALUES ($1)", team.Name)
	if err != nil {
		return err
	}
	res, err := a.db.Query("SELECT id FROM users WHERE is_admin = true")
	if err != nil {
		return err
	}
	for res.Next() {
		var user database.User
		res.Scan(&user.Id)
		_, err := a.db.Query("INSERT INTO teams_users (team_id, user_id, role_id) VALUES ($1, $2, 1)", team.Id, user.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Admin) AddUserToTeam(userId int, teamId int, roleId int) error {
	_, err := a.db.Query("INSERT INTO users_teams (user_id, team_id, role_id) VALUES ($1, $2, $3)", userId, teamId, roleId)
	return err
}

func (a Admin) GetTeamMembers(teamId int) ([]database.User, error) {
	users := []database.User{}
	res, err := a.db.Query("SELECT id, username, name, surname, phone, email, is_admin FROM users WHERE id IN (SELECT user_id FROM users_teams WHERE team_id = $1);", teamId)
	res.Scan(&users)
	for res.Next() {
		var user database.User
		res.Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Phone, &user.Email, &user.IsAdmin)
		users = append(users, user)
	}

	return users, err
}

func (a Admin) CanEditTeamUser(userId int, teamId int) (bool, error) {
	var canEdit bool
	err := a.db.QueryRow("SELECT can_edit_users FROM roles WHERE id IN (SELECT role_id FROM users_teams WHERE user_id = $1 AND team_id = $2)", userId, teamId).Scan(&canEdit)
	return canEdit, err
}

func (a Admin) CanEditTeamProject(userId int, teamId int) (bool, error) {
	var canEdit bool
	err := a.db.QueryRow("SELECT can_edit_projects FROM roles WHERE id IN (SELECT role_id FROM users_teams WHERE user_id = $1 AND team_id = $2)", userId, teamId).Scan(&canEdit)
	return canEdit, err
}
