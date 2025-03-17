package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
			ID:        user.GetID(),
			Status:    user.GetStatus(),
			Name:      user.GetName(),
			Email:     user.GetEmail(),
			Gender:    user.GetGender(),
			BirthDate: user.GetBirthDate(),
		}
	}

	c.JSON(http.StatusOK, res)
}

type listUserRes struct {
	ID        string        `json:"id" binding:"required"`
	Status    domain.Status `json:"status" binding:"required"`
	Name      string        `json:"name" binding:"required"`
	Email     string        `json:"email" binding:"required"`
	Gender    domain.Gender `json:"gender" binding:"required"`
	BirthDate time.Time     `json:"birthDate" binding:"required"`
}
