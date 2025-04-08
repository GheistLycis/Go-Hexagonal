package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func list(c *gin.Context, s domain.UserServicePort) {
	filters := domain.ListUsersServiceFiltersDTO{}

	if nameQuery, nameExists := c.GetQuery("name"); nameExists {
		filters.Name = &nameQuery
	}
	if statusQuery, statusExists := c.GetQuery("status"); statusExists {
		status := domain.Status(statusQuery)
		filters.Status = &status
	}
	if genderQuery, genderExists := c.GetQuery("gender"); genderExists {
		gender := domain.Gender(genderQuery)
		filters.Gender = &gender
	}

	users, err := s.List(filters)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	res := make([]listUserRes, len(users))

	for i, user := range users {
		res[i] = listUserRes{
			ID:        user.ID,
			Status:    user.Status,
			Name:      user.Name,
			Email:     user.Email,
			Gender:    user.Gender,
			BirthDate: user.BirthDate,
		}
	}

	c.JSON(http.StatusOK, res)
}

type listUserRes struct {
	ID        uuid.UUID     `json:"id" binding:"required"`
	Status    domain.Status `json:"status" binding:"required"`
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
	BirthDate time.Time     `json:"birthDate" binding:"required"`
}
