package api

import (
	app "Go-Hexagonal/app/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func enable(c *gin.Context, s app.UserServiceI) {
	id := c.Param("id")
	updatedBy := c.GetHeader("userEmail") // TODO: session service with auth (JWT)

	if _, err := s.Enable(id, updatedBy); err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
