package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-backend-1/entities"
)

func AuthorizationWithRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		actor, ok := c.MustGet("actor").(entities.Actor)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		found := false
		for _, role := range roles {
			if role == actor.Role.RoleName {
				found = true
				break
			}
		}
		if !found {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
