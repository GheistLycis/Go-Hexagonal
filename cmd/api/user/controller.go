package api

import (
	app "Go-Hexagonal/app/user"
	db "Go-Hexagonal/infra/db/user"

	"github.com/gin-gonic/gin"
)

var service app.UserServiceI

func init() {
	repo := db.NewUserRepo()
	service = app.NewUserService(repo)
}

func SetRouter(g *gin.Engine) {
	g.GET("user/:id", handle(getById))
}

func handle(m method) gin.HandlerFunc {
	return func(c *gin.Context) {
		m(c, service)
	}
}

type method func(c *gin.Context, s app.UserServiceI)
