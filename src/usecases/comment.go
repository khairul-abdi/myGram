package usecases

import (
	"fmt"
	"myGram/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *uc) GetComments(c *gin.Context, userID int) ([]models.CommentGetResponse, string, int) {
	var ListComment []models.CommentGetResponse

	whereVariable := "user_id = ?"
	whereValue := userID
	res, err := u.repo.FindAllComment(whereVariable, whereValue)
	if err != nil {
		return nil, "error while get findAll data comment", http.StatusBadRequest
	}

	for _, v := range res {
		fmt.Println("ISI V ===>>> ", v)
		ListComment = append(ListComment, v.ToCommentGetResponse())
	}

	return ListComment, "Success", http.StatusOK
}
