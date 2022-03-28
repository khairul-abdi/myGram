package models

import "time"

type Photo struct {
	Id        int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title     string    `gorm:"type:varchar(100);column:title;not null;" valid:"Required"`
	Caption   string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl  string    `gorm:"type:varchar(100);column:photo_url;not null;" valid:"Required"`
	UserId    int       `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      User      `gorm:"foreignKey:UserId"`
}

type PhotoGetResponse struct {
	Id        int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title     string    `gorm:"type:varchar(100);column:title"`
	Caption   string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl  string    `gorm:"type:varchar(100);column:photo_url"`
	UserId    int       `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      UserUpdateRequest
}

type PhotoResponse struct {
	Id        int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoRequest struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title" valid:"Required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"Required"`
	UserId   int    `json:"user_id"`
}

func (r Photo) ToPhotoGetResponse() PhotoGetResponse {
	return PhotoGetResponse{
		Id:        r.Id,
		Title:     r.Title,
		Caption:   r.Caption,
		PhotoUrl:  r.PhotoUrl,
		UserId:    r.UserId,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		User:      r.User.ToUserUpdateRequest(),
	}
}

func (r Photo) ToPhotoResponse() PhotoResponse {
	return PhotoResponse{
		Id:        r.Id,
		Title:     r.Title,
		Caption:   r.Caption,
		PhotoUrl:  r.PhotoUrl,
		UserId:    r.UserId,
		CreatedAt: r.CreatedAt,
	}
}

func (r Photo) ToPhotoRequest() PhotoRequest {
	return PhotoRequest{
		Id:       r.Id,
		Title:    r.Title,
		Caption:  r.Caption,
		PhotoUrl: r.PhotoUrl,
		UserId:   r.UserId,
	}
}
