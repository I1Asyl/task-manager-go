package repositories

import (
	"database/sql"
	"fmt"

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
	err = a.db.QueryRow("SELECT id FROM teams WHERE name = $1", team.Name).Scan(&team.Id)
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
		fmt.Println(team.Id, user.Id)
		_, err := a.db.Query("INSERT INTO users_teams (team_id, user_id, role_id) VALUES ($1, $2, 1)", team.Id, user.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Admin) DeleteTeam(teamId int) error {
	_, err := a.db.Query("DELETE FROM users_teams WHERE team_id = $1", teamId)
	if err != nil {
		return err
	}
	_, err = a.db.Query("DELETE FROM teams WHERE id = $1", teamId)
	if err != nil {
		return err
	}

	return nil
}
