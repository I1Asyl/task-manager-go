package repositories

import (
	"database/sql"
	"errors"

	"github.com/I1Asyl/task-manager-go/database"
)

type Authorization struct {
	db *sql.DB
}

func NewAuthorization(db *sql.DB) *Authorization {
	return &Authorization{
		db: db,
	}
}

func (auth Authorization) CreateUser(user database.User) error {
	_, err := auth.db.Query("INSERT INTO users (username, password, name, surname, phone, email, is_admin) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.Username, user.Password, user.Name, user.Surname, user.Phone, user.Email, user.IsAdmin)
	return err
}

func (aut Authorization) CreateTeam(team database.Team) error {
	_, err := aut.db.Query("INSERT INTO teams (name) VALUES ($1)", team.Name)
	return err
}

func (auth Authorization) GetUser(user database.UserForm) (database.User, error) {
	var userReturn database.User
	var err error
	err = auth.db.QueryRow("SELECT id, username, name, surname, is_admin, password FROM users WHERE username = $1;", user.Username).Scan(&userReturn.Id, &userReturn.Username, &userReturn.Name, &userReturn.Surname, &userReturn.IsAdmin, &userReturn.Password)

	if err != nil {
		return database.User{}, err
	}
	if user.Password != userReturn.Password {
		return database.User{}, errors.New("no password")
	}

	return userReturn, nil
}
func (auth Authorization) AddSession(session database.Session) error {
	_, err := auth.db.Query("INSERT INTO sessions (user_id, first_token, token) VALUES ($1, $2, $3);", session.UserId, session.FirstToken, session.Token)
	return err
}

func (auth Authorization) CheckRefreshToken(first_token, token string) (bool, error) {
	var token_res string
	err := auth.db.QueryRow("SELECT token FROM sessions WHERE first_token = $1;", first_token).Scan(&token_res)
	if err != nil {
		return false, err
	} else if token_res != token {
		return false, errors.New("xxx")
	}
	return true, nil
}

func (auth Authorization) GetUserByFirstToken(first_token string) (database.User, error) {
	var user_id int
	err := auth.db.QueryRow("SELECT user_id FROM sessions WHERE first_token = $1;", first_token).Scan(&user_id)
	if err != nil {
		return database.User{}, err
	}
	var user database.User
	err = auth.db.QueryRow("SELECT id, username, name, surname, is_admin FROM users WHERE id = $1;", user_id).Scan(&user.Id, &user.Username, &user.Name, &user.Surname, &user.IsAdmin)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}

func (auth Authorization) DeleteToken(first_token string) error {
	_, err := auth.db.Query("DELETE FROM sessions WHERE first_token = $1;", first_token)
	return err
}

func (auth Authorization) UpdateToken(first_token string, token string) error {
	_, err := auth.db.Query("UPDATE sessions SET token = $1 WHERE first_token = $2;", token, first_token)
	return err
}
