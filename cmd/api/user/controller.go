package api

import (
	app "Go-Hexagonal/app/user"
	db "Go-Hexagonal/infra/db"
	db_user "Go-Hexagonal/infra/db/user"

	"github.com/gin-gonic/gin"
)

/*
SetRouter maps all routes in User context to their handlers.

-g: gin server
*/
func SetRouter(g *gin.Engine) {
	repo := db_user.NewUserRepo(db.DB) // TODO: create deps container
	service := app.NewUserService(repo)

	g.GET("user/:id", handle(getById, service))
	g.GET("user", handle(list, service))
	g.POST("user", handle(create, service))
	g.POST("user/:id/enable", handle(enable, service))
	g.POST("user/:id/disable", handle(disable, service))
}

func handle(m method, s app.UserServiceI) gin.HandlerFunc {
	return func(c *gin.Context) {
		m(c, s)
	}
}

type method func(c *gin.Context, s app.UserServiceI)
