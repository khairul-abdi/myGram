package controllers

import (
	"myGram/packages"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (ctrl *ctrl) GetSocialMedias(c *gin.Context) {
	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Idint := int(userID)
	res, message, code := ctrl.uc.GetSocialMedias(c, Idint)

	packages.ResponseSocial(c, message, code, res)
}

func (ctrl *ctrl) StoreSocialMedia(c *gin.Context) {

	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	data, message, code := ctrl.uc.StoreSocialMedia(c, userID)

	packages.Response(c, message, code, data)
}

func (ctrl *ctrl) UpdateSocialMedia(c *gin.Context) {
	//Get socialmediaId from param
	socialmediaIdint, err := strconv.Atoi(c.Param("socialmediaId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	res, message, code := ctrl.uc.UpdateSocialMedia(c, socialmediaIdint)

	packages.Response(c, message, code, res)
}

func (ctrl *ctrl) DeleteSocialMedia(c *gin.Context) {
	socialmediaIdint, err := strconv.Atoi(c.Param("socialmediaId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	err = ctrl.uc.DeleteSocialMedia(c, socialmediaIdint)
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
		return
	}

	message := "Your social media has been successfully deleted"
	code := http.StatusOK

	packages.Response(c, message, code, nil)

}
