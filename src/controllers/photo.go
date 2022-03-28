package controllers

import (
	"myGram/packages"

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

func (ctrl *ctrl) StorePhotos(c *gin.Context) {
	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	data, message, code := ctrl.uc.StorePhotos(c, userID)

	packages.Response(c, message, code, data)
}
