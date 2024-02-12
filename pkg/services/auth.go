package services

import (
	"errors"
	"strings"
	"time"

	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/repositories"
	jwt "github.com/dgrijalva/jwt-go"
)

type Authorization struct {
	repo *repositories.Repository
}

func NewAuthorization(repo *repositories.Repository) *Authorization {
	return &Authorization{repo: repo}
}

func (a Authorization) CreateUser(model database.Model) (map[string]string, error) {

	user := database.User(model.User)
	if mistakes := user.IsValid(); len(mistakes) > 0 {
		return mistakes, nil
	}
	err := a.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (a Authorization) CreateTeam(model database.Model) (map[string]string, error) {

	team := database.Team(model.Team)
	if mistakes := team.IsValid(); len(mistakes) > 0 {
		return mistakes, nil
	}
	err := a.repo.CreateTeam(team)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func generateTokens(user database.User, firstToken string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.Id,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Unix() + 60*60,
	})
	accessTokenString, err := accessToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}
	if firstToken == "" {
		firstToken = accessTokenString
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"first_token": firstToken,
		"exp":         time.Now().Unix() + 60*60*24*30,
	})
	refreshTokenTokenString, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenTokenString, nil
}

func refreshToDB(token string) string {
	refreshTokenSplit := strings.Split(token, ".")
	refreshToken := refreshTokenSplit[2]
	return refreshToken
}

func (a Authorization) RefreshToken(tokenString string) (string, string, error) {
	claims, err := parseToken(tokenString)
	if err != nil {
		return "", "", err
	}
	firstToken := claims["first_token"].(string)

	valid, err := a.repo.CheckRefreshToken(firstToken, refreshToDB(tokenString))
	if err != nil && err.Error() == "xxx" {
		a.repo.DeleteToken(firstToken)
	}
	if err != nil || !valid {

		return "", "", err
	}
	user, err := a.repo.GetUserByFirstToken(firstToken)
	if err != nil {
		return "", "", err
	}
	access, refresh, err := generateTokens(user, firstToken)
	if err != nil {
		return "", "", err
	}
	err = a.repo.UpdateToken(firstToken, refreshToDB(refresh))
	if err != nil {
		return "", "", err
	}
	return access, refresh, err
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
func (a Authorization) VerifyUser(tokenString string) bool {
	_, err := parseToken(tokenString)
	if err != nil {
		return false
	}
	return true
}

func (a Authorization) GetUserIdByToken(token string) (int, error) {
	claims, err := parseToken(token)
	if err != nil {
		return 0, err
	}
	valF := claims["user_id"].(float64)

	val := int(valF)
	return val, nil
}
func (a Authorization) VerifyAdmin(tokenString string) bool {
	claims, err := parseToken(tokenString)
	if err != nil {
		return false
	}
	if _, ok := claims["is_admin"]; !ok {
		return false
	}
	if claim, ok := claims["is_admin"].(bool); ok && claim {
		return true
	}
	return false
}

func (a Authorization) Logout(token string) error {
	claims, err := parseToken(token)
	if err != nil {
		return err
	}
	firstToken, ok := claims["first_token"]
	if !ok {
		return errors.New("Token is invalid")
	}
	firstTokenString, ok := firstToken.(string)
	if !ok {
		return errors.New("Token is invalid")
	}
	return a.repo.DeleteToken(firstTokenString)
}

func (a Authorization) Login(model database.Model) (string, string, map[string]string, error) {

	userForm := database.UserForm(model.UserForm)
	if mistakes := userForm.IsValid(); len(mistakes) > 0 {
		return "", "", mistakes, nil
	}

	var user database.User
	var err error
	if user, err = a.repo.GetUser(userForm); err != nil || (user == database.User{}) {

		return "", "", nil, err
	}

	access, refresh, er := generateTokens(user, "")

	if er != nil {
		return "", "", nil, er
	}

	a.repo.AddSession(database.Session{UserId: user.Id, FirstToken: access, Token: refreshToDB(refresh)})
	return access, refresh, nil, nil
}
