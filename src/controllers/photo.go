package controllers

import (
	"myGram/packages"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (ctrl *ctrl) GetPhotos(c *gin.Context) {
	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Idint := int(userID)
	res, message, code := ctrl.uc.GetPhotos(c, Idint)

	packages.ResponsePhoto(c, message, code, res)

}

func (ctrl *ctrl) StorePhoto(c *gin.Context) {
	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	data, message, code := ctrl.uc.StorePhotos(c, userID)

	packages.Response(c, message, code, data)
}

func (ctrl *ctrl) UpdatePhoto(c *gin.Context) {
	//Get photoId from param
	photoIdInt, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	res, message, code := ctrl.uc.UpdatePhotos(c, photoIdInt)

	packages.Response(c, message, code, res)
}

func (ctrl *ctrl) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	err = ctrl.uc.DeletePhoto(c, photoId)
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
		return
	}

	message := "Your photo has been successfully deleted"
	code := http.StatusOK

	packages.Response(c, message, code, nil)
}
