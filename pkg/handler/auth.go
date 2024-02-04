package handler

import (
	"github.com/I1Asyl/task-manager-go/pkg/services"
)

type Auth struct {
	services services.Service
}

func NewAuth(services services.Service) *Auth {
	return &Auth{services: services}
}
