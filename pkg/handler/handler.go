package handler

import (
	"net/http"

	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services services.Service
}

// New returns a new instance of a gin server
func New(services services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Assign() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", h.getStatus)

	return r
}

func (h *Handler) getStatus(c *gin.Context) {
	c.String(http.StatusOK, h.services.ReturnStatus())
}
