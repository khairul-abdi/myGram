package usecases

import (
	"myGram/packages"
	"myGram/src/models"
	"net/http"
	"time"

	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
)

func (u *uc) GetComments(c *gin.Context, userId int) ([]models.CommentGetResponse, string, int) {
	var ListComment []models.CommentGetResponse

	whereVariable := "user_id = ?"
	whereValue := userId
	res, err := u.repo.FindAllComment(whereVariable, whereValue)
	if err != nil {
		return nil, "error while get findAll data comment", http.StatusBadRequest
	}

	for _, v := range res {
		ListComment = append(ListComment, v.ToCommentGetResponse())
	}

	return ListComment, "Success", http.StatusOK
}

func (u *uc) StoreComment(c *gin.Context, userId int) (map[string]interface{}, string, int) {
	var comment models.CommentRequest
	c.ShouldBindJSON(&comment)

	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&comment)
	if err != nil {
		return nil, "Internal Server Error", http.StatusInternalServerError
	}

	if !b {
		// validation does not pass
		var message string
		for i, err := range valid.Errors {
			if i == 0 {
				message = err.Message
			} else {
				message += ", " + err.Message
			}
		}
		return nil, message, http.StatusBadRequest
	}

	comment.UserId = userId

	//store comment
	data := comment.ToComment()
	err = u.repo.InsertOneComment(&data)
	if err != nil {
		code, message := packages.Validate(err)
		return nil, message, code
	}

	res := map[string]interface{}{
		"id":         data.Id,
		"message":    data.Message,
		"photo_id":   data.PhotoId,
		"user_id":    data.UserId,
		"created_at": time.Now(),
	}

	return res, "Success create comment", 200

}

func (u *uc) UpdateComment(c *gin.Context, commentId int) (map[string]interface{}, string, int) {
	var comment models.CommentRequest
	c.ShouldBindJSON(&comment)

	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&comment)
	if err != nil {
		return nil, "Internal Server Error", http.StatusInternalServerError
	}

	if !b {
		// validation does not pass
		var message string
		for i, err := range valid.Errors {
			if i == 0 {
				message = err.Message
			} else {
				message += ", " + err.Message
			}
		}
		return nil, message, http.StatusBadRequest
	}

	data := comment.ToComment()
	data.Id = commentId
	res, err := u.repo.UpdateComment(&data)

	if err != nil {
		return nil, "Internal Server Error", http.StatusInternalServerError
	}

	result := map[string]interface{}{
		"id":         res.Id,
		"message":    res.Message,
		"title":      res.Photo.Title,
		"caption":    res.Photo.Caption,
		"photo_url":  res.Photo.PhotoUrl,
		"user_id":    res.UserId,
		"updated_at": res.UpdatedAt,
	}

	return result, "Success update comment", http.StatusOK
}
