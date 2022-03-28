package middlewares

import (
	"myGram/packages"
	"myGram/src/usecases"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeSocialMedia() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get socialmediaId from param
		socialMediaId, err := strconv.Atoi(c.Param("socialmediaId"))
		if err != nil {
			message := "invalid parameter"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}

		// Get UserId
		UserData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(UserData["id"].(float64))

		//Get UserId from socmed
		socMed, err := usecases.GetSocialMediaById(c, socialMediaId)
		if err != nil {
			message := "social media not found"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}
		if int(userID) != socMed.UserId {
			message := "Access Forbidden"
			packages.Response(c, message, http.StatusForbidden, nil)
			return
		}
	}
}
