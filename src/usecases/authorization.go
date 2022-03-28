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
