package repositories

import (
	"myGram/src/models"
	"time"
)

func (r *repo) UpdateUser(oldUser, entity *models.User) error {
	err := r.db.Model(&oldUser).Updates(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateSocialMedia(OldSocialMedia, NewSocialMedia *models.SocialMedia) error {
	err := r.db.Model(&OldSocialMedia).Updates(&NewSocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdatePhoto(OldPhoto, NewPhoto *models.Photo) error {
	err := r.db.Model(&OldPhoto).Updates(&NewPhoto).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) UpdateComment(data *models.Comment) (*models.Comment, error) {
	var entity models.Comment
	entity.Message = data.Message

	r.db.First(&data)
	data.Message = entity.Message
	data.UpdatedAt = time.Now()

	err := r.db.Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
