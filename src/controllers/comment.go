package controllers

import (
	"myGram/packages"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (ctrl *ctrl) GetComments(c *gin.Context) {
	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	res, message, code := ctrl.uc.GetComments(c, userID)

	packages.ResponseComment(c, message, code, res)
}
