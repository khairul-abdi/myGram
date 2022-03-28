package controllers

import (
	"myGram/src/usecases"

	"github.com/gin-gonic/gin"
)

type ctrl struct {
	uc usecases.UC
}

type Controllers interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	GetPhotos(c *gin.Context)
	StorePhoto(c *gin.Context)
	UpdatePhoto(c *gin.Context)
	DeletePhoto(c *gin.Context)

	GetSocialMedias(c *gin.Context)
	StoreSocialMedia(c *gin.Context)
	UpdateSocialMedia(c *gin.Context)
	DeleteSocialMedia(c *gin.Context)
}

func NewCtrl(uc usecases.UC) Controllers {
	return &ctrl{uc: uc}
}
