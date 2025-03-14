package api

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(server *gin.Engine) {
	server.GET("user/:id", getById)
}