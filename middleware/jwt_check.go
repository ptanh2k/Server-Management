package middleware

import (
	"net/http"
	t "sm/pkg/token"

	"github.com/gin-gonic/gin"
)

func JwtCheckMiddleware() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		err := t.ValidToken(c)

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}

	return gin.HandlerFunc(fn)
}
