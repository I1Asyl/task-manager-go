package database

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

type Model struct {
	CurrentUser User
	User        User     `json:"user`
	UserForm    UserForm `json:"user_form" db:"user_form"`
	Team        Team     `json:"team" db:"team"`
	Role        Role     `json:"role"`
	Project     Project  `json:"project"`
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin"`
}

type Team struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UserForm struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
type Session struct {
	Id         int    `json:"id" db:"id"`
	UserId     int    `json:"user_id" db:"user_id"`
	FirstToken string `json:"first_token" db:"first_token"`
	Token      string `json:"token" db:"token"`
}

type Role struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type UserTeam struct {
	UserId int `json:"user_id" db:"user_id"`
	TeamId int `json:"team_id" db:"team_id"`
	RoleId int `json:"role_id" db:"role_id"`
}

type Project struct {
	Id            int    `json:"id" db:"XX"`
	Name          string `json:"name" db:"XXXX"`
	Description   string `json:"description" db:"XXXXXXXXXXX"`
	TeamId        int    `json:"team_id" db:"XXXXX"`
	CurrentStatus string `json:"current_status" db:"XXXXXX"`
}

func (u UserForm) IsValid() map[string]string {
	errors := make(map[string]string)
	if err := validUsername(u.Username); err != nil {
		errors["username"] = err.Error()
	}
	if err := validPassword(u.Password); err != nil {
		errors["password"] = err.Error()
	}
	return errors
}
func (t Team) IsValid() map[string]string {
	errors := make(map[string]string)
	if err := validName(t.Name); err != nil {
		errors["name"] = err.Error()
	}
	return errors
}

func (u User) IsValid() map[string]string {
	errors := make(map[string]string)
	if err := validName(u.Name); err != nil {
		errors["name"] = err.Error()
	}
	if err := validName(u.Username); err != nil {
		errors["username"] = err.Error()
	}
	if err := validName(u.Surname); err != nil {
		errors["surname"] = err.Error()
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		errors["email"] = err.Error()
	}
	if err := validEmail(u.Email); err != nil {
		errors["email"] = err.Error()
	}
	if err := validPhone(u.Phone); err != nil {
		errors["phone"] = err.Error()
	}
	if err := validPassword(u.Password); err != nil {
		errors["password"] = err.Error()
	}
	return errors
}

func validUsername(username string) error {
	if len(username) < 3 {
		return errors.New("Username is too short")
	}
	return nil
}
func validName(name string) error {
	if len(name) < 3 {
		return errors.New("Name is too short")
	}
	return nil
}
func validEmail(name string) error {
	if len(name) < 3 {
		return errors.New("Name is too short")
	}
	return nil
}
func validPhone(phone_number string) error {
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phone_number = strings.ReplaceAll(phone_number, " ", "")

	if !re.MatchString(phone_number) {
		return errors.New("Invalid phone number")
	}
	return nil
}

func validPassword(password string) error {
	if len(password) < 6 {
		return errors.New("Password is too short")
	}
	passRegex := `\d`
	re := regexp.MustCompile(passRegex)
	if !re.MatchString(password) {
		return errors.New("Password has no digits")
	}
	passRegex = `[A-Za-z]`
	re = regexp.MustCompile(passRegex)
	if !re.MatchString(password) {
		return errors.New("Password has no letters")
	}

	return nil
}
