package middlewares

import (
	"myGram/packages"
	"myGram/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			message := "Unauthenticated"
			code := http.StatusUnauthorized
			packages.Response(c, message, code, nil)
			return
		}
		c.Set("userData", verifyToken)

		c.Next()
	}
}
