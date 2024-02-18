package repositories

import (
	"database/sql"

	"github.com/I1Asyl/task-manager-go/database"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db}
}

func (a User) AddUserToTeam(userId int, teamId int, roleId int) error {
	_, err := a.db.Query("INSERT INTO users_teams (user_id, team_id, role_id) VALUES ($1, $2, $3)", userId, teamId, roleId)
	return err
}

func (a User) GetTeamMembers(teamId int) ([]database.User, error) {
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
