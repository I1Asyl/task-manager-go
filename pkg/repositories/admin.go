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
	return err
}

func (a Admin) AddUserToTeam(user database.User, team database.Team) error {
	_, err := a.db.Query("INSERT INTO users_teams (team_id, user_id) VALUES ($1, $2)", team.Id, user.Id)
	return err
}

func (a Admin) GetTeamMembers(team database.Team) ([]database.User, error) {
	users := []database.User{}
	res, err := a.db.Query("SELECT id, username, name, surname, phone, email, is_admin FROM users WHERE id IN (SELECT user_id FROM users_teams WHERE team_id = $1);", team.Id)
	res.Scan(&users)
	for res.Next() {
		var user database.User
		res.Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.Phone, &user.Email, &user.IsAdmin)
		users = append(users, user)
	}

	return users, err
}
