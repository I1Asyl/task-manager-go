package handler

import (
	"errors"

	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type User struct {
	services services.Service
}

func NewUser(services services.Service) *User {
	return &User{services: services}
}

func (a User) addUserToTeam(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	var user database.Model
	user.CurrentUser.Id = userId.(int)

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := a.services.AddUserToTeam(user); err != nil {
		ctx.AbortWithError(400, err)
	}
	ctx.JSON(200, user)
}

func (a User) getTeamMembers(ctx *gin.Context) {
	var team database.Model
	if err := ctx.BindJSON(&team); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	ans, err := a.services.GetTeamMembers(team)
	if err != nil {
		ctx.AbortWithError(400, err)
	}
	ctx.JSON(200, ans)
}

func (a User) checkUser(ctx *gin.Context) {
	_, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("user id does not exist"))
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

func (a User) logout(ctx *gin.Context) {
	var user struct {
		Token string `json:"token"`
	}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Given data is not corect"})
		return
	}
	if err := a.services.Logout(user.Token); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
