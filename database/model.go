package database

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

type Model struct {
	User     User     `json:"user"`
	UserForm UserForm `json:"user_form"`
	Team     Team     `json:"team"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	TeamId   int    `json:"team_id"`
	IsAdmin  bool   `json:"is_admin"`
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Session struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	FirstToken string `json:"first_token"`
	Token      string `json:"token"`
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
