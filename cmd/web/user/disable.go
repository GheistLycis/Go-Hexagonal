package web

import (
	app_ports "Go-Hexagonal/app/user/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

func disable(c *gin.Context, s app_ports.UserServicePort) {
	id := c.Param("id")
	updatedBy := c.GetHeader("userEmail")

	if _, err := s.Disable(id, updatedBy); err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
