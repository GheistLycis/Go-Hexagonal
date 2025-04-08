package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getById(c *gin.Context, s domain.UserServicePort) {
	id := c.Param("id")

	user, err := s.Get(domain.GetUserServiceFiltersDTO{ID: &id})
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	res := getUserByIdRes{
		ID:        user.ID,
		Status:    user.Status,
		Name:      user.Name,
		Email:     user.Email,
		Gender:    user.Gender,
		BirthDate: user.BirthDate,
	}

	c.JSON(http.StatusOK, res)
}

type getUserByIdRes struct {
	ID        uuid.UUID     `json:"id" binding:"required"`
	Status    domain.Status `json:"status" binding:"required"`
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
	BirthDate time.Time     `json:"birthDate" binding:"required"`
}
