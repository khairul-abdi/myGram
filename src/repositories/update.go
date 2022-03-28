package repositories

import "myGram/src/models"

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
