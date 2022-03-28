package repositories

import (
	"myGram/src/models"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	InsertOne(data interface{}) error
	FindOne(data interface{}, whereVariable string, whereValue interface{}) error
	FindAll(whereVariable string, whereValue interface{}) ([]models.SocialMedia, error)

	UpdateSocialMedia(OldSocialMedia, NewSocialMedia *models.SocialMedia) error
	DeleteSocialMedia(id int, SocialMedia *models.SocialMedia) error

	UpdateUser(oldUser, entity *models.User) error
	DeleteUser(id int, User *models.User) error
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
