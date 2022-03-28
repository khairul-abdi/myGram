package usecases

import (
	"myGram/packages"
	"myGram/src/models"
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
)

func (u *uc) GetSocialMedias(c *gin.Context, userId int) ([]models.SocialMediaGetResponse, string, int) {
	var cList []models.SocialMediaGetResponse

	whereVariable := "user_id = ?"
	whereValue := userId
	res, err := u.repo.FindAll(whereVariable, whereValue)
	if err != nil {
		return nil, "error while get findAll data social_media", http.StatusBadRequest
	}

	for _, v := range res {
		cList = append(cList, v.ToSocialMediaGetResponse())
	}

	return cList, "Success", http.StatusOK
}

func (u *uc) StoreSocialMedia(c *gin.Context, userId int) (map[string]interface{}, string, int) {
	SocialMedia := models.SocialMedia{}
	c.ShouldBindJSON(&SocialMedia)

	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&SocialMedia)
	if err != nil {
		return nil, "Internal Server Error", http.StatusInternalServerError
	}

	if !b {
		// validation does not pass
		var message string
		for i, err := range valid.Errors {
			if i == 0 {
				message = err.Message
			} else {
				message += ", " + err.Message
			}
		}
		return nil, message, http.StatusBadRequest
	}
	SocialMedia.UserId = uint(userId)
	err = u.repo.InsertOne(&SocialMedia)
	if err != nil {
		code, message := packages.Validate(err)
		return nil, message, code
	}

	data := map[string]interface{}{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserId,
		"created_at":       SocialMedia.CreatedAt,
	}

	return data, "Success create photo", 200
}

func (u *uc) UpdateSocialMedia(c *gin.Context, socialmediaIdint int) (map[string]interface{}, string, int) {
	NewSocialMedia := models.SocialMedia{}
	OldSocialMedia := models.SocialMedia{}

	c.ShouldBindJSON(&NewSocialMedia)

	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&NewSocialMedia)
	if err != nil {
		return nil, "Internal Server Error", http.StatusInternalServerError
	}

	if !b {
		// validation does not pass
		var message string
		for i, err := range valid.Errors {
			if i == 0 {
				message = err.Message
			} else {
				message += ", " + err.Message
			}
		}
		return nil, message, http.StatusBadRequest
	}

	whereVariable := "id = ?"
	whereValue := socialmediaIdint

	err = u.repo.FindOne(&OldSocialMedia, whereVariable, whereValue)
	if err != nil {
		return nil, "Data Not Found", http.StatusBadRequest
	}

	err = u.repo.UpdateSocialMedia(&OldSocialMedia, &NewSocialMedia)
	if err != nil {
		return nil, "Bad Request", http.StatusBadRequest
	}

	data := map[string]interface{}{
		"id":               OldSocialMedia.ID,
		"name":             OldSocialMedia.Name,
		"social_media_url": OldSocialMedia.SocialMediaUrl,
		"user_id":          OldSocialMedia.UserId,
		"created_at":       OldSocialMedia.UpdatedAt,
	}
	return data, "Success update social media", http.StatusOK
}

func (u *uc) DeleteSocialMedia(c *gin.Context, socialmediaIdint int) error {
	SocialMedia := models.SocialMedia{}
	err := u.repo.DeleteSocialMedia(socialmediaIdint, &SocialMedia)
	if err != nil {
		return err
	}

	return nil
}
