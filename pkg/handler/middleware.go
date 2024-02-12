package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (h Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		headerParts := strings.Split(header, " ")
		if headerParts[0] != "Bearer" || len(headerParts) != 2 || !h.auth.verifyUser(headerParts[1]) {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}

		id, err := h.auth.getUserIdByToken(headerParts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		}
		c.Set("userId", id)

		c.Next()
	}
}

func (h Handler) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerParts := strings.Split(header, " ")
		if headerParts[0] != "Bearer" || len(headerParts) != 2 || !h.auth.verifyAdmin(headerParts[1]) {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
			return
		}
		id, err := h.auth.getUserIdByToken(headerParts[1])
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
		}
		c.Set("userId", id)
		c.Set("token", headerParts[1])
		c.Next()
	}
}
