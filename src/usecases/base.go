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
	UpdateUser(c *gin.Context, userId uint) (map[string]interface{}, string, int)
	DeleteUser(c *gin.Context, userId int) error

	GetPhotos(c *gin.Context, photoId int) ([]models.PhotoGetResponse, string, int)
	StorePhotos(c *gin.Context, userId uint) (map[string]interface{}, string, int)
	UpdatePhotos(c *gin.Context, photoId int) (map[string]interface{}, string, int)
	DeletePhoto(c *gin.Context, photoId int) error

	GetComments(c *gin.Context, userId int) ([]models.CommentGetResponse, string, int)
	StoreComment(c *gin.Context, userId int) (map[string]interface{}, string, int)
	UpdateComment(c *gin.Context, commentId int) (map[string]interface{}, string, int)
	DeleteComment(c *gin.Context, commentId int) error

	GetSocialMedias(c *gin.Context, userId int) ([]models.SocialMediaGetResponse, string, int)
	StoreSocialMedia(c *gin.Context, userId int) (map[string]interface{}, string, int)
	UpdateSocialMedia(c *gin.Context, socialmediaIdint int) (map[string]interface{}, string, int)
	DeleteSocialMedia(c *gin.Context, socialmediaIdint int) error
}

func NewUC(repo repositories.Repo) UC {
	return &uc{repo: repo}
}
