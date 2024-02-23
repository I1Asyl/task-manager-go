package handler

import (
	"errors"
	"strconv"

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

// addUserToTeam godoc
// @Summary      Add user
// @Description  Add new user to the team.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        model body database.Model true "Enter user id, team id and role id"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /userTeam [post]
func (a User) addUserToTeam(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	var user database.Model
	user.CurrentUser.Id = userId.(int)

	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Error with input data"})
		return
	}

	if err := a.services.AddUserToTeam(user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not add user to the team"})
		return
	}
	ctx.JSON(200, user)
}

// getTeamMembers godoc
// @Summary      Get team members
// @Description  Get all users in the taeam
// @Tags         user
// @Produce      json
// @Param        team_id  path int true "Enter team id"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /userTeam/{team_id} [get]
func (a User) getTeamMembers(ctx *gin.Context) {
	var team database.Model
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	team.CurrentUser.Id = userId.(int)
	temp, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	team.Team.Id = temp
	ans, err := a.services.GetTeamMembers(team)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not get team members"})
		return
	}
	ctx.JSON(200, ans)
}

// createProject godoc
// @Summary      Create project
// @Description  Create a project and assign it to the team
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        model body database.Model true "Enter project info and team id"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /project [post]
func (a User) createProject(ctx *gin.Context) {
	var project database.Model
	if err := ctx.BindJSON(&project); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Error with input data"})
		return
	}
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	project.CurrentUser.Id = userId.(int)
	if err := a.services.CreateProject(project); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not create a project"})
		return
	}
	ctx.JSON(200, project)
}

// checkUser godoc
// @Summary      Check user
// @Description  Chech if user exists based on authorization header
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      401  {object}  string
// @Router       /check [get]
func (a User) checkUser(ctx *gin.Context) {
	_, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("user id does not exist"))
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

// logout godoc
// @Summary      log out
// @Description  Deactivate given refresh token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        model body token true "Enter refresh token"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /logout [post]
func (a User) logout(ctx *gin.Context) {
	var user token
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Error with input data"})
		return
	}
	if err := a.services.Logout(user.Token); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not log out"})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

// logout godoc
// @Summary      create a task
// @Description  create a task and assign it to someone and project.
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        model body database.Model true "Enter task info and team id"
// @Param        Authorization header string  true  "Authorization header"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      401  {object}  string
// @Failure      422  {object}  string
// @Router       /task [post]
func (a User) createTask(ctx *gin.Context) {
	var task database.Model
	if err := ctx.BindJSON(&task); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Error with input data"})
		return
	}
	userId, exists := ctx.Get("userId")
	if !exists {
		panic(errors.New("no user id"))
	}
	task.CurrentUser.Id = userId.(int)
	if err := a.services.CreateTask(task); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(422, gin.H{"message": "Could not create a task"})
		return
	}
	ctx.JSON(200, task)
}
