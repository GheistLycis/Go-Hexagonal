package user

import (
	app "Go-Hexagonal/src/user/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func disable(c *gin.Context, s app.UserServicePort) {
	id := c.Param("id")
	updatedBy := c.GetHeader("userEmail")

	if _, err := s.Disable(id, updatedBy); err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
