package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	user "Go-Hexagonal/src/user/cmd/web"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(DB *gorm.DB) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[WEB] Recovered from panic -", r)
		}
	}()

	server := gin.Default()
	serverPort, err := strconv.ParseInt(os.Getenv("WEB_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("[WEB] Failed to parse ENV variable WEB_PORT - %v", err)
	}

	initRouters(server, DB)
	initHandlers(server)
	server.Run(fmt.Sprintf(":%d", serverPort))
}

func initRouters(s *gin.Engine, DB *gorm.DB) {
	user.SetRouter(s, DB)
}

func initHandlers(s *gin.Engine) {
	s.Use(handleErrors)
}

func handleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err

		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
