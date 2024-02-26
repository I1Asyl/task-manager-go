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
	Task        Task     `json:"task"`
	UserTeam    UserTeam `json:"user_team"`
}

type Task struct {
	Id            int    `json:"id" db:"XX"`
	UserId        int    `json:"user_id" db:"XXXXX"`
	AssignerId    int    `json:"assigner_id" db:"XXXXX"`
	StartTime     string `json:"start_time" db:"XXXXX"`
	EndTime       string `json:"end_time" db:"XXXXX"`
	Name          string `json:"name" db:"XXXX"`
	Description   string `json:"description" db:"XXXXXXXXXXX"`
	ProjectId     int    `json:"project_id" db:"XXXXXXXXXX"`
	CurrentStatus string `json:"current_status" db:"XXXXXX"`
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
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"XXXXXXXXXXX"`
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
	Id            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Description   string `json:"description" db:"description"`
	TeamId        int    `json:"team_id" db:"XXXXX"`
	CurrentStatus string `json:"current_status" db:"XXXXXX"`
}

func isValidForUpdate(allNames []string, allValues []interface{}, allFunctions []func(interface{}) error) (map[string]string, error) {
	names := []string{}
	values := []interface{}{}
	functions := []func(interface{}) error{}
	for i := 0; i < len(allValues); i++ {
		if allValues[i] != nil {
			names = append(names, allNames[i])
			values = append(values, allValues[i])
			functions = append(functions, allFunctions[i])
		}
	}
	return isValidForInsert(names, values, functions)

}

func isValidForInsert(names []string, values []interface{}, functions []func(interface{}) error) (map[string]string, error) {
	mistakes := make(map[string]string)
	if len(names) != len(values) || len(names) != len(functions) {
		return nil, errors.New("invalid number of names, values and functions in validation")
	}
	for i := 0; i < len(names); i++ {
		if err := functions[i](values[i]); err != nil {
			mistakes[names[i]] = err.Error()
		}
	}
	return mistakes, nil
}

func isValid(names []string, values []interface{}, functions []func(interface{}) error, validityCode int) map[string]string {
	if validityCode == 1 {
		mistakes, err := isValidForInsert(names, values, functions)
		if err != nil {
			panic(err)
		}
		return mistakes
	}
	if validityCode == 2 {
		mistakes, err := isValidForUpdate(names, values, functions)
		if err != nil {
			panic(err)
		}
		return mistakes
	}
	return nil
}

func (t Task) IsValid(validityCode int) map[string]string {
	names := []string{"name", "description", "project_id", "assigner_id"}
	values := []interface{}{t.Name, t.Description, t.ProjectId, t.AssignerId}
	functions := []func(interface{}) error{validName, validName, validId, validId}
	return isValid(names, values, functions, validityCode)
}

func (u UserForm) IsValid(validityCode int) map[string]string {
	names := []string{"username", "password"}
	values := []interface{}{u.Username, u.Password}
	functions := []func(interface{}) error{validUsername, validPassword}
	return isValid(names, values, functions, validityCode)
}

func (p Project) IsValid(validityCode int) map[string]string {
	names := []string{"name", "description", "team_id"}
	values := []interface{}{p.Name, p.Description, p.TeamId}
	functions := []func(interface{}) error{validName, validName, validId}
	return isValid(names, values, functions, validityCode)
}
func (t Team) IsValid(validityCode int) map[string]string {
	names := []string{"name"}
	values := []interface{}{t.Name}
	functions := []func(interface{}) error{validName}
	return isValid(names, values, functions, validityCode)
}

func (u User) IsValid(validityCode int) map[string]string {
	names := []string{"name", "username", "surname", "email", "phone", "password"}
	values := []interface{}{u.Name, u.Username, u.Surname, u.Email, u.Phone, u.Password}
	functions := []func(interface{}) error{validName, validUsername, validName, validEmail, validPhone, validPassword}
	return isValid(names, values, functions, validityCode)
}

func validUsername(input interface{}) error {
	username := input.(string)
	if len(username) < 3 {
		return errors.New("username is too short")
	}
	return nil
}
func validName(input interface{}) error {
	name := input.(string)
	if len(name) < 3 {
		return errors.New("name is too short")
	}
	return nil
}
func validEmail(input interface{}) error {
	email := input.(string)
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}
	return nil
}
func validPhone(input interface{}) error {
	phone_number := input.(string)
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phone_number = strings.ReplaceAll(phone_number, " ", "")

	if !re.MatchString(phone_number) {
		return errors.New("invalid phone number")
	}
	return nil
}

func validId(input interface{}) error {
	id := input.(int)
	if id <= 0 {
		return errors.New("id is not given or negative")
	}
	return nil
}

func validPassword(input interface{}) error {
	password := input.(string)
	if len(password) < 6 {
		return errors.New("password is too short")
	}
	passRegex := `\d`
	re := regexp.MustCompile(passRegex)
	if !re.MatchString(password) {
		return errors.New("password has no digits")
	}
	passRegex = `[A-Za-z]`
	re = regexp.MustCompile(passRegex)
	if !re.MatchString(password) {
		return errors.New("password has no letters")
	}

	return nil
}
