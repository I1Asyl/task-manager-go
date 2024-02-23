package handler

import (
	"errors"
	"strconv"

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

// createUser godoc
// @Summary      Create an user
// @Description  Create an user from json file with user as a key and structure as a value.
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        user body database.Model  true  "User information"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /user [post]
func (a Admin) createUser(ctx *gin.Context) {
	var user database.Model
	//var us database.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if mistakes, err := a.services.CreateUser(user); len(mistakes) > 0 || err != nil {
		if err != nil {
			ctx.Error(err)
		}
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not login", "errors": mistakes})
		return
	}
	ctx.JSON(200, gin.H{"message": "User created"})
}

// createTeam godoc
// @Summary      Create a team
// @Description  Create an team from json file with team as a key and structure as a value.
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        team body database.Model  true  "Team information"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /team [post]
func (a Admin) createTeam(ctx *gin.Context) {
	var team database.Model
	_, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	if err := ctx.BindJSON(&team); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if mistakes, err := a.services.CreateTeam(team); len(mistakes) > 0 || err != nil {
		if err != nil {
			ctx.Error(err)
		}
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not login", "errors": mistakes})
		return
	}
	ctx.JSON(200, gin.H{"message": "Team created"})
}

// createTeam godoc
// @Summary      Dealete a team
// @Description  Delete a team by its id.
// @Tags         admin
// @Produce      json
// @Param        team_id path int  true  "Team id"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      500  {object}  string
// @Router       /team/{team_id} [delete]
func (a Admin) deleteTeam(ctx *gin.Context) {
	model := database.Model{}
	var err error
	model.Team.Id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	if err := a.services.DeleteTeam(model); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Could not delete team"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Team deleted"})
}
