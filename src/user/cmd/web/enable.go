package user

import (
	domain "Go-Hexagonal/src/user/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func enable(c *gin.Context, s domain.UserServicePort) {
	id := c.Param("id")
	updatedBy := c.GetHeader("userEmail")

	if _, err := s.Enable(id, updatedBy); err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
