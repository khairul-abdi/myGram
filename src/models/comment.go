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
