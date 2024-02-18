package handler

import (
	"errors"

	"github.com/I1Asyl/task-manager-go/database"
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	services services.Service
}

func NewAdmin(service services.Service) *Admin {
	return &Admin{services: service}
}

func (a Admin) createUser(ctx *gin.Context) {
	var user database.Model
	//var us database.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if mistakes, err := a.services.CreateUser(user); len(mistakes) > 0 {
		if err != nil {
			ctx.Error(err)
		}
		ctx.AbortWithStatusJSON(400, mistakes)
		return
	}
	ctx.JSON(200, user)
}

func (a Admin) createTeam(ctx *gin.Context) {
	var team database.Model
	_, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	if err := ctx.BindJSON(&team); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if mistakes, err := a.services.CreateTeam(team); len(mistakes) > 0 {
		if err != nil {
			ctx.Error(err)
		}
		ctx.AbortWithStatusJSON(400, mistakes)
		return
	}
	ctx.JSON(200, team)
}
