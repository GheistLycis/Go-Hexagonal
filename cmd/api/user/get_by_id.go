package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getById(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}