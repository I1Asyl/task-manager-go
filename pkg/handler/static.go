package handler

import (
	"net/http"

	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type Static struct {
	services services.Service
}

func NewStatic(services services.Service) *Static {
	return &Static{services: services}
}

func (s Static) getStatus(c *gin.Context) {
	c.String(http.StatusOK, s.services.ReturnStatus())
}
