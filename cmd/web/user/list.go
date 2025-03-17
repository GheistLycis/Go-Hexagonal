package web

import (
	app_ports "Go-Hexagonal/app/user/ports"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func list(c *gin.Context, s app_ports.UserServicePort) {
	filters := app_ports.ListUsersServiceFiltersDTO{}

	if nameQuery, nameExists := c.GetQuery("name"); nameExists {
		filters.Name = &nameQuery
	}
	if statusQuery, statusExists := c.GetQuery("status"); statusExists {
		status := app_ports.Status(statusQuery)
		filters.Status = &status
	}
	if genderQuery, genderExists := c.GetQuery("gender"); genderExists {
		gender := app_ports.Gender(genderQuery)
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
	ID        string           `json:"id" binding:"required"`
	Status    app_ports.Status `json:"status" binding:"required"`
	Name      string           `json:"name" binding:"required"`
	Email     string           `json:"email" binding:"required"`
	Gender    app_ports.Gender `json:"gender" binding:"required"`
	BirthDate time.Time        `json:"birthDate" binding:"required"`
}
