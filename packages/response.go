package packages

import (
	"myGram/src/models"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, message string, code int, data map[string]interface{}) {
	if len(data) != 0 {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
			"data":    data,
		})
	} else {
		c.AbortWithStatusJSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func ResponseSocial(c *gin.Context, message string, code int, data []models.SocialMediaGetResponse) {
	if len(data) != 0 {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
			"data":    data,
		})
	} else {
		c.AbortWithStatusJSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func ResponsePhoto(c *gin.Context, message string, code int, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
