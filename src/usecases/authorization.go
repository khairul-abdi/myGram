package usecases

import (
	"myGram/src/models"
	"myGram/src/repositories"

	"github.com/gin-gonic/gin"
)

func GetSocialMediaById(c *gin.Context, socialMediaId int) (*models.SocialMediaResponse, error) {

	whereVariable := "id= ?"
	socialMedia, err := repositories.FindOneSocialMediaById(whereVariable, socialMediaId)
	if err != nil {
		return nil, err
	}

	socmed := socialMedia.ToSocialMediaResponse()

	return &socmed, nil

}

func GetPhotoById(c *gin.Context, photoId int) (*models.PhotoResponse, error) {

	whereVariable := "id= ?"
	photo, err := repositories.FindOnePhotoById(whereVariable, photoId)
	if err != nil {
		return nil, err
	}

	entity := photo.ToPhotoResponse()

	return &entity, nil
}
