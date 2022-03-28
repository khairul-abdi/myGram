package usecases

import (
	"myGram/src/models"
	"myGram/src/repositories"

	"github.com/gin-gonic/gin"
)

type uc struct {
	repo repositories.Repo
}

type UC interface {
	UserRegister(c *gin.Context) (map[string]interface{}, string, int)
	UserLogin(c *gin.Context) (map[string]interface{}, string, int)
	UpdateUser(c *gin.Context, userID uint) (map[string]interface{}, string, int)
	DeleteUser(c *gin.Context, userID int) error

	GetSocialMedias(c *gin.Context, UserId int) ([]models.SocialMediaGetResponse, string, int)
	StoreSocialMedia(c *gin.Context, UserId int) (map[string]interface{}, string, int)
	UpdateSocialMedia(c *gin.Context, socialmediaIdint int) (map[string]interface{}, string, int)
	DeleteSocialMedia(c *gin.Context, socialmediaIdint int) error
}

func NewUC(repo repositories.Repo) UC {
	return &uc{repo: repo}
}
