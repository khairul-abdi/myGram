package models

import "time"

type Comment struct {
	Id        int       `gorm:"primarykey;autoIncrement:true;column:id"`
	UserId    int       `gorm:"column:user_id"`
	PhotoId   int       `gorm:"column:photo_id"`
	Message   string    `gorm:"type:varchar(100);column:message"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      User      `gorm:"foreignKey:UserId"`
	Photo     Photo     `gorm:"foreignKey:PhotoId"`
}

type CommentRequest struct {
	Id      int    `json:"id,omitempty"`
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message" valid:"Required"`
}

type CommentGetResponse struct {
	Id        int       `gorm:"primarykey;autoIncrement:true;column:id"`
	UserId    int       `gorm:"column:user_id"`
	PhotoId   int       `gorm:"column:photo_id"`
	Message   string    `gorm:"type:varchar(100);column:message"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	User      UserUpdateRequest
	Photo     PhotoRequest
}

type CommentResponse struct {
	Id        int       `gorm:"primarykey;column:id;autoIncrement:true"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (r Comment) ToCommentResponse() CommentResponse {
	return CommentResponse{
		Id:        r.Id,
		UserId:    r.UserId,
		PhotoId:   r.PhotoId,
		Message:   r.Message,
		CreatedAt: r.CreatedAt,
	}
}

func (r Comment) ToCommentGetResponse() CommentGetResponse {
	return CommentGetResponse{
		Id:        r.Id,
		Message:   r.Message,
		PhotoId:   r.PhotoId,
		UserId:    r.UserId,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		User:      r.User.ToUserUpdateRequest(),
		Photo:     r.Photo.ToPhotoRequest(),
	}
}
