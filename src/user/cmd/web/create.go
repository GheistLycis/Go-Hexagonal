package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func create(c *gin.Context, s domain.UserServicePort) {
	payload := domain.CreateUserServiceDTO{}
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
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}

	c.JSON(http.StatusOK, res)
}

type createUserRes struct {
	ID        uuid.UUID     `json:"id" binding:"required"`
	Status    domain.Status `json:"status" binding:"required"`
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
	BirthDate time.Time     `json:"birthDate" binding:"required"`
}
