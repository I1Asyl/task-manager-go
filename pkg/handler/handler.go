package handler

import (
	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type auth interface {
}
type static interface {
	getStatus(c *gin.Context)
}

type Handler struct {
	auth
	static
}

// New returns a new instance of a gin server
func New(services *services.Service) *Handler {
	return &Handler{auth: NewAuth(*services), static: NewStatic(*services)}
}

func (h *Handler) Assign() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", h.static.getStatus)

	return r
}
