package usecases

import (
	"myGram/packages"
	"myGram/src/helpers"
	"myGram/src/models"
	"net/http"
	"time"

	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
)

func (u *uc) UserRegister(c *gin.Context) (map[string]interface{}, string, int) {
	User := models.User{}
	c.ShouldBindJSON(&User)

	// Form Validate
	valid := validation.Validation{}
	b, err := valid.Valid(&User)
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

	err = u.repo.InsertOne(&User)
	if err != nil {
		code, message := packages.Validate(err)
		return nil, message, code
	}

	data := map[string]interface{}{
		"id":       User.Id,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	}

	return data, "Success insert data", http.StatusCreated
}

func (u *uc) UserLogin(c *gin.Context) (map[string]interface{}, string, int) {
	User := models.User{}
	c.ShouldBindJSON(&User)

	password := User.Password
	whereVariable := "email = ?"
	whereValue := User.Email

	err := u.repo.FindOne(&User, whereVariable, whereValue)
	if err != nil {
		return nil, "invalid email", http.StatusUnauthorized
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		return nil, "Wrong password", http.StatusUnauthorized
	}

	data := map[string]interface{}{
		"token": helpers.GenerateToken(User.Id, User.Email),
	}

	return data, "Success Create Token", http.StatusOK
}

func (u *uc) UpdateUser(c *gin.Context, userID uint) (map[string]interface{}, string, int) {

	var (
		newUser models.UserUpdateRequest
		oldUser models.User
		entity  models.User
	)
	c.ShouldBindJSON(&newUser)
	valid := validation.Validation{}
	b, err := valid.Valid(&newUser)
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

	whereVariable := "id = ?"
	whereValue := userID
	err = u.repo.FindOne(&oldUser, whereVariable, whereValue)
	if err != nil {
		return nil, "Data Not Found", http.StatusBadRequest
	}

	entity.Id = userID
	entity.Username = newUser.Username
	entity.Email = newUser.Email
	entity.Password = oldUser.Password
	entity.Age = oldUser.Age
	entity.CreatedAt = oldUser.CreatedAt
	entity.UpdatedAt = time.Now()

	err = u.repo.UpdateUser(&oldUser, &entity)
	if err != nil {
		return nil, "Bad Request", http.StatusBadRequest
	}

	data := map[string]interface{}{
		"id":         oldUser.Id,
		"email":      oldUser.Email,
		"username":   oldUser.Username,
		"age":        oldUser.Age,
		"updated_at": oldUser.UpdatedAt,
	}

	return data, "Success Update user", http.StatusOK
}

func (u *uc) DeleteUser(c *gin.Context, userID int) error {
	User := models.User{}
	err := u.repo.DeleteUser(userID, &User)
	if err != nil {
		return err
	}

	return nil
}
