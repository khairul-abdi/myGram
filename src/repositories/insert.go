package repositories

import "myGram/src/models"

func (r *repo) InsertOne(data interface{}) error {
	err := r.db.Debug().Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) InsertOneComment(data *models.Comment) error {
	err := r.db.Debug().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}
