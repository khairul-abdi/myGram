package models

import "time"

type Photo struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title      string    `gorm:"type:varchar(100);column:title;not null;" valid:"Required"`
	Caption    string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl   string    `gorm:"type:varchar(100);column:photo_url;not null;" valid:"Required"`
	UserId     int       `gorm:"column:user_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       User      `gorm:"foreignKey:UserId"`
}

type PhotoGetResponse struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title      string    `gorm:"type:varchar(100);column:title"`
	Caption    string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl   string    `gorm:"type:varchar(100);column:photo_url"`
	UserId     int       `gorm:"column:user_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       UserUpdateRequest
}

func (r Photo) ToPhotoGetResponse() PhotoGetResponse {
	return PhotoGetResponse{
		Id:         r.Id,
		Title:      r.Title,
		Caption:    r.Caption,
		PhotoUrl:   r.PhotoUrl,
		UserId:     r.UserId,
		CreatedAt:  r.CreatedAt,
		Updated_At: r.Updated_At,
		User:       r.User.ToUserUpdateRequest(),
	}
}
