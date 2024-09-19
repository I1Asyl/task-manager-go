package handler

import (
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type auth interface {
	login(ctx *gin.Context)
	main(ctx *gin.Context)
	refreshToken(ctx *gin.Context)
}

type admin interface {
	createUser(ctx *gin.Context)
	createTeam(ctx *gin.Context)
	deleteTeam(ctx *gin.Context)
}
type user interface {
	addUserToTeam(ctx *gin.Context)
	getTeamMembers(ctx *gin.Context)
	logout(ctx *gin.Context)
	checkUser(ctx *gin.Context)
	createProject(ctx *gin.Context)
	createTask(ctx *gin.Context)
	getTasks(ctx *gin.Context)
	getTasksByProject(ctx *gin.Context)
	updateTask(ctx *gin.Context)
	updateProject(ctx *gin.Context)
	getProjects(ctx *gin.Context)
}

type middleware interface {
	verifyUser(jwt string) bool
	verifyAdmin(jwt string) bool
	getUserIdByToken(token string) (int, error)
	UserMiddleware() gin.HandlerFunc
	AdminMiddleware() gin.HandlerFunc
}

type Handler struct {
	middleware
	auth
	user
	admin
}

// New returns a new instance of a gin server
func New(services *services.Service) *Handler {
	return &Handler{auth: NewAuth(*services), user: NewUser(*services), admin: NewAdmin(*services), middleware: NewMiddleware(*services)}
}

func (h Handler) Assign() *gin.Engine {
	router := gin.New()
	router.Use(logger.SetLogger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.POST("/login", h.auth.login)
		api.GET("", h.auth.main)
		api.POST("/refresh", h.auth.refreshToken)
		authorized := api.Group("user")
		{
			authorized.Use(h.UserMiddleware())
			authorized.POST("/logout", h.user.logout)
			authorized.GET("/check", h.user.checkUser)

			authorized.POST("/team/user", h.user.addUserToTeam)
			authorized.GET("/team/:id/user", h.user.getTeamMembers)

			authorized.POST("/task", h.user.createTask)
			authorized.GET("/task", h.user.getTasks)
			authorized.PUT("/task", h.user.updateTask)

			authorized.POST("/project", h.user.createProject)
			authorized.PUT("/project", h.user.updateProject)
			authorized.GET("/project", h.user.getProjects)

			authorized.GET("/project/:id/task", h.user.getTasksByProject)
		}

		admin := api.Group("admin")
		{
			admin.Use(h.AdminMiddleware())
			admin.POST("/user", h.createUser)
			admin.POST("/team", h.createTeam)
			admin.DELETE("/team/:id", h.deleteTeam)

		}
	}

	return router
}
