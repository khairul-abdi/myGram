package usecases

import (
	"myGram/packages"
	"myGram/src/models"
	"net/http"

	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
)

func (u *uc) GetPhotos(c *gin.Context, photoId int) ([]models.PhotoGetResponse, string, int) {
	var photoList []models.PhotoGetResponse

	whereVariable := "user_id = ?"
	whereValue := photoId
	res, err := u.repo.FindAllPhoto(whereVariable, whereValue)
	if err != nil {
		return nil, "error while get findAll data photos", http.StatusBadRequest
	}

	for _, v := range res {
		photoList = append(photoList, v.ToPhotoGetResponse())
	}

	return photoList, "Success", http.StatusOK
}

func (u *uc) StorePhotos(c *gin.Context, userId uint) (map[string]interface{}, string, int) {
	var photo models.Photo
	c.ShouldBindJSON(&photo)
	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&photo)
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

	photo.UserId = int(userId)
	err = u.repo.InsertOne(&photo)
	if err != nil {
		code, message := packages.Validate(err)
		return nil, message, code
	}

	data := map[string]interface{}{
		"id":         photo.Id,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserId,
		"created_at": photo.CreatedAt,
	}

	return data, "Success create photo", 200
}

func (u *uc) UpdatePhotos(c *gin.Context, photoId int) (map[string]interface{}, string, int) {
	NewPhoto := models.Photo{}
	OldPhoto := models.Photo{}

	c.ShouldBindJSON(&NewPhoto)
	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&NewPhoto)
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
	whereValue := photoId

	err = u.repo.FindOne(&OldPhoto, whereVariable, whereValue)
	if err != nil {
		return nil, "Data Not Found", http.StatusBadRequest
	}

	err = u.repo.UpdatePhoto(&OldPhoto, &NewPhoto)
	if err != nil {
		return nil, "Bad Request", http.StatusBadRequest
	}

	data := map[string]interface{}{
		"id":         OldPhoto.Id,
		"title":      OldPhoto.Title,
		"caption":    OldPhoto.Caption,
		"photo_url":  OldPhoto.PhotoUrl,
		"user_id":    OldPhoto.UserId,
		"updated_at": OldPhoto.UpdatedAt,
	}

	return data, "Success update photo", http.StatusOK
}

func (u *uc) DeletePhoto(c *gin.Context, photoId int) error {
	Photo := models.Photo{}
	err := u.repo.DeletePhoto(photoId, &Photo)
	if err != nil {
		return err
	}

	return nil
}
