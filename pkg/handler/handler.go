package handler

import (
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type auth interface {
	login(ctx *gin.Context)
	verifyUser(jwt string) bool
	verifyAdmin(jwt string) bool
	refreshToken(ctx *gin.Context)
	getUserIdByToken(token string) (int, error)
	logout(ctx *gin.Context)
	checkUser(ctx *gin.Context)
}

type admin interface {
	createUser(ctx *gin.Context)
	createTeam(ctx *gin.Context)
	addUserToTeam(ctx *gin.Context)
}
type static interface {
	getStatus(c *gin.Context)
}

type Handler struct {
	auth
	static
	admin
}

// New returns a new instance of a gin server
func New(services *services.Service) *Handler {
	return &Handler{auth: NewAuth(*services), static: NewStatic(*services), admin: NewAdmin(*services)}
}

func (h Handler) Assign() *gin.Engine {
	router := gin.New()
	router.Use(logger.SetLogger())
	router.Use(gin.Recovery())
	router.GET("/ping", h.static.getStatus)

	router.GET("/login", h.auth.login)
	router.GET("/refresh", h.auth.refreshToken)
	authorized := router.Group("")
	{
		authorized.Use(h.AuthMiddleware())
		authorized.POST("/logout", h.logout)
		authorized.GET("/check", h.checkUser)

	}

	admin := router.Group("")
	{
		admin.Use(h.AdminMiddleware())
		admin.POST("/user", h.createUser)
		admin.POST("/team", h.createTeam)
		admin.POST("/teamUser", h.addUserToTeam)

	}

	return router
}
