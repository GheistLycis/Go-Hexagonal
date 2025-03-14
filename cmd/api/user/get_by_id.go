package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUserByIdRes struct {
	ID     string  `json:"id" binding:"required"`
	Status string  `json:"status" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	Email  string  `json:"email" binding:"required"`
}

func getById(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}