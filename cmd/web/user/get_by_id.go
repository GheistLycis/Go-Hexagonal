package web

import (
	app "Go-Hexagonal/app/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getById(c *gin.Context, s app.UserServiceI) {
	id := c.Param("id")

	user, err := s.Get(app.GetUserServiceFilters{ID: &id})
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	res := getUserByIdRes{
		ID:        user.GetID(),
		Status:    user.GetStatus(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		Gender:    user.GetGender(),
		BirthDate: user.GetBirthDate(),
	}

	c.JSON(http.StatusOK, res)
}

type getUserByIdRes struct {
	ID        string     `json:"id" binding:"required"`
	Status    app.Status `json:"status" binding:"required"`
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Gender    app.Gender `json:"gender" binding:"required"`
	BirthDate time.Time  `json:"birthDate" binding:"required"`
}
