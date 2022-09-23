package middelware

import (
	"auth/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Request.Header["Token"]; !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "missing token")
			return
		}

		token := c.Request.Header["Token"][0]

		valid, err := jwt.ValidateToken(token, "V3ry_S3cr3t_C0D3")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
			return
		}
		c.Next()
	}
}
