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

func (r *repo) FindAll(whereVariable string, whereValue interface{}) ([]models.SocialMedia, error) {
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

func FindOneSocialMediaById(whereVariable string, socialMediaId int) (*models.SocialMedia, error) {
	var entity models.SocialMedia
	db := database.GetDB()
	err := db.Debug().First(&entity, "id =?", socialMediaId).Error

	if err != nil {
		return nil, err
	}
	return &entity, nil
}
