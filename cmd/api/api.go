package api

import (
	"fmt"
	"net/http"

	user_controller "Go-Hexagonal/cmd/api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
Init creates the server, set its routers and handlers and then runs it.

-p: the server port.
-DB: the database connection
*/
func Init(p int, DB *gorm.DB) {
	server := gin.Default()

	initRouters(server, DB)
	initHandlers(server)
	server.Run(fmt.Sprintf(":%d", p))
}

func initRouters(s *gin.Engine, DB *gorm.DB) {
	user_controller.SetRouter(s, DB)
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
