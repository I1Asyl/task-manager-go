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

	router.GET("/login", h.auth.login)
	router.GET("/refresh", h.auth.refreshToken)
	authorized := router.Group("")
	{
		authorized.Use(h.UserMiddleware())
		authorized.POST("/logout", h.user.logout)
		authorized.GET("/check", h.user.checkUser)

		authorized.POST("/teamUser", h.user.addUserToTeam)
		authorized.GET("/teamUser", h.user.getTeamMembers)

		authorized.POST("/project", h.user.createProject)
	}

	admin := router.Group("")
	{
		admin.Use(h.AdminMiddleware())
		admin.POST("/user", h.createUser)
		admin.POST("/team", h.createTeam)
		admin.DELETE("/team/:id", h.deleteTeam)

	}

	return router
}
