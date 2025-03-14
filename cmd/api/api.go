package api

import (
	"fmt"
	"net/http"

	user_controller "Go-Hexagonal/cmd/api/user"

	"github.com/gin-gonic/gin"
)

/*
Init creates the server, set its routers and handlers and then runs it.

-port: the server port.
*/
func Init(port int) {
	server := gin.Default()

	initRouters(server)
	initHandlers(server)
	server.Run(fmt.Sprintf(":%d", port))
}

func initRouters(server *gin.Engine) {
	user_controller.SetRouter(server)
}

func initHandlers(server *gin.Engine) {
	server.Use(handleErrors)
}

func handleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err

		c.JSON(http.StatusInternalServerError, err.Error())
	}
}