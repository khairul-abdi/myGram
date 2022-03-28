package models

import (
	"time"
)

type SocialMedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null;column:username" valid:"Required"`
	SocialMediaUrl string    `gorm:"type:text;not null;column:social_media_url" valid:"Required"`
	UserId         uint      `gorm:"column:user_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	User           User
}

type SocialMediaGetResponse struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null;column:username;" valid:"Required"`
	SocialMediaUrl string    `gorm:"type:text;not null;column:social_media_url;" valid:"Required"`
	UserId         uint      `gorm:"column:user_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	User           UserUpdateRequest
}

type SocialMediaResponse struct {
	Id             int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

func (r SocialMedia) ToSocialMediaGetResponse() SocialMediaGetResponse {
	return SocialMediaGetResponse{
		ID:             r.ID,
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserId:         r.UserId,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
		User:           r.User.ToUserUpdateRequest(),
	}
}

func (r SocialMedia) ToSocialMediaResponse() SocialMediaResponse {
	return SocialMediaResponse{
		Id:             int(r.ID),
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserId:         int(r.UserId),
		CreatedAt:      r.CreatedAt,
	}
}
