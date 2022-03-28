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

	FindAllPhoto(whereVariable string, whereValue interface{}) ([]models.Photo, error)
	UpdatePhoto(OldPhoto, NewPhoto *models.Photo) error
	DeletePhoto(id int, photo *models.Photo) error

	FindAllSocialMedia(whereVariable string, whereValue interface{}) ([]models.SocialMedia, error)
	UpdateSocialMedia(OldSocialMedia, NewSocialMedia *models.SocialMedia) error
	DeleteSocialMedia(id int, SocialMedia *models.SocialMedia) error

	FindAllComment(whereVariable string, whereValue interface{}) ([]models.Comment, error)

	UpdateUser(oldUser, entity *models.User) error
	DeleteUser(id int, User *models.User) error
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}
