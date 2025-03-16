package api

import (
	db_user "Go-Hexagonal/adapters/db/user"
	app "Go-Hexagonal/app/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
SetRouter maps all routes in User context to their handlers.

-g: the gin server
-DB: the database active connection
*/
func SetRouter(g *gin.Engine, DB *gorm.DB) {
	repo := db_user.NewUserRepo(DB) // TODO: implement singleton deps container (or not?)
	service := app.NewUserService(repo)

	g.GET("user/:id", handle(getById, service))
	g.GET("user", handle(list, service))
	g.POST("user", handle(create, service))
	g.POST("user/:id/enable", handle(enable, service))
	g.POST("user/:id/disable", handle(disable, service))
}

func handle(m method, s app.UserServiceI) gin.HandlerFunc { // TODO: implement auth service
	return func(c *gin.Context) {
		m(c, s)
	}
}

type method func(c *gin.Context, s app.UserServiceI)
