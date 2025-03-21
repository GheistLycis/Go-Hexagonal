package user

import (
	app "Go-Hexagonal/src/user/app"
	domain "Go-Hexagonal/src/user/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context, s app.UserServicePort) {
	payload := app.CreateUserServiceDTO{}
	createdBy := c.GetHeader("userEmail")

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.Create(payload, createdBy)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	res := createUserRes{
		ID:        user.GetID(),
		Status:    user.GetStatus(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		Gender:    user.GetGender(),
		BirthDate: user.GetBirthDate(),
	}

	c.JSON(http.StatusOK, res)
}

type createUserRes struct {
	ID        string        `json:"id" binding:"required"`
	Status    domain.Status `json:"status" binding:"required"`
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
	BirthDate time.Time     `json:"birthDate" binding:"required"`
}
