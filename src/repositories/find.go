package repositories

import (
	"myGram/packages/database"
	"myGram/src/models"

	"gorm.io/gorm/clause"
)

func (r *repo) FindOne(data interface{}, whereVariable string, whereValue interface{}) error {
	err := r.db.Debug().Where(whereVariable, whereValue).Take(data).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) FindAllSocialMedia(whereVariable string, whereValue interface{}) ([]models.SocialMedia, error) {
	var List []models.SocialMedia
	err := r.db.Preload(clause.Associations).Find(&List, whereVariable, whereValue).Error
	if err != nil {
		return nil, err
	}
	return List, nil
}

func (r *repo) FindAllPhoto(whereVariable string, whereValue interface{}) ([]models.Photo, error) {
	var List []models.Photo
	err := r.db.Preload(clause.Associations).Find(&List, whereVariable, whereValue).Error
	if err != nil {
		return nil, err
	}

	return List, nil
}

func (r *repo) FindAllComment(whereVariable string, whereValue interface{}) ([]models.Comment, error) {
	var ListComment []models.Comment
	err := r.db.Preload(clause.Associations).Find(&ListComment, whereVariable, whereValue).Error
	if err != nil {
		return nil, err
	}

	return ListComment, nil
}

func FindOneSocialMediaById(whereVariable string, socialMediaId int) (*models.SocialMedia, error) {
	var entity models.SocialMedia
	db := database.GetDB()
	err := db.Debug().First(&entity, "id =?", socialMediaId).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func FindOnePhotoById(whereVariable string, photoId int) (*models.Photo, error) {
	var entity models.Photo
	db := database.GetDB()
	err := db.Debug().First(&entity, "id =?", photoId).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func FindOneCommentById(whereVariable string, commentId int) (*models.Comment, error) {
	var entity models.Comment
	db := database.GetDB()
	err := db.Debug().First(&entity, "id =?", commentId).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}
