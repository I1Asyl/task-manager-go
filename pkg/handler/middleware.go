package handler

import (
	"strings"

	"github.com/I1Asyl/task-manager-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	services services.Service
}

func NewMiddleware(services services.Service) *Middleware {
	return &Middleware{services: services}
}

func (a Middleware) UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		headerParts := strings.Split(header, " ")
		if headerParts[0] != "Bearer" || len(headerParts) != 2 || !a.verifyUser(headerParts[1]) {
			c.AbortWithStatusJSON(401, gin.H{"message": "unauthorized"})
			return
		}

		id, err := a.getUserIdByToken(headerParts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		}
		c.Set("userId", id)

		c.Next()
	}
}

func (a Middleware) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerParts := strings.Split(header, " ")
		if headerParts[0] != "Bearer" || len(headerParts) != 2 || !a.verifyAdmin(headerParts[1]) {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		id, err := a.getUserIdByToken(headerParts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		}
		c.Set("userId", id)
		c.Set("token", headerParts[1])
		c.Next()
	}
}

func (a Middleware) verifyUser(jwt string) bool {
	return a.services.VerifyUser(jwt)
}

func (a Middleware) verifyAdmin(jwt string) bool {
	return a.services.VerifyAdmin(jwt)
}

func (a Middleware) getUserIdByToken(token string) (int, error) {
	return a.services.GetUserIdByToken(token)
}
