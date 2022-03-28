package repositories

import (
	"myGram/src/models"
)

func (r *repo) DeleteSocialMedia(id int, SocialMedia *models.SocialMedia) error {

	err := r.db.Delete(&SocialMedia, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteUser(id int, User *models.User) error {
	err := r.db.Delete(&User, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeletePhoto(id int, photo *models.Photo) error {
	err := r.db.Delete(&photo, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) DeleteComment(id int, Comment *models.Comment) error {

	err := r.db.Delete(&Comment, id).Error
	if err != nil {
		return err
	}
	return nil
}
