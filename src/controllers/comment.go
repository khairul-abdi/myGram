package controllers

import (
	"myGram/packages"
	"net/http"
	"strconv"

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

func (ctrl *ctrl) StoreComment(c *gin.Context) {
	// Get userId
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	data, message, code := ctrl.uc.StoreComment(c, userID)

	packages.Response(c, message, code, data)
}

func (ctrl *ctrl) UpdateComment(c *gin.Context) {
	//Get commentId from param
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		message := "Internal Server Error"
		code := http.StatusInternalServerError
		packages.Response(c, message, code, nil)
	}

	res, message, code := ctrl.uc.UpdateComment(c, commentId)

	packages.Response(c, message, code, res)
}
