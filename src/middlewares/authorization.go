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

		// Get userId
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		//Get userId from socmed
		socMed, err := usecases.GetSocialMediaById(c, socialMediaId)
		if err != nil {
			message := "social media not found"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}

		// compare
		if userID != socMed.UserId {
			message := "Access Forbidden"
			packages.Response(c, message, http.StatusForbidden, nil)
			return
		}
	}
}

func AuthorizePhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get photoId from param
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			message := "invalid parameter"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}

		// Get userId
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		//Get userId from photo
		photo, err := usecases.GetPhotoById(c, photoId)
		if err != nil {
			message := "photo not found"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}

		//compare
		if userID != photo.UserId {
			message := "Access Forbidden"
			packages.Response(c, message, http.StatusForbidden, nil)
			return
		}
	}
}

func AuthorizeComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get commentId from param
		commentId, err := strconv.Atoi(c.Param("commentId"))
		if err != nil {
			message := "invalid parameter"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}

		// Get userId
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		//Get userId from comment
		comment, err := usecases.GetCommentById(c, commentId)
		if err != nil {
			message := "comment not found"
			packages.Response(c, message, http.StatusInternalServerError, nil)
			return
		}

		//compare
		if userID != comment.UserId {
			message := "Access Forbidden"
			packages.Response(c, message, http.StatusForbidden, nil)
			return
		}
	}
}
