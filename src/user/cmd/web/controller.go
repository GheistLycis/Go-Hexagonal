package user

import (
	postgres_adapters "Go-Hexagonal/infra/postgres/adapters"
	app "Go-Hexagonal/src/user/app"
	infra "Go-Hexagonal/src/user/infra"
	ports "Go-Hexagonal/src/user/ports"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
SetRouter maps all routes in User context to their handlers.

-g: the gin server
-db: the database active connection
*/
func SetRouter(g *gin.Engine, db *gorm.DB) {
	dbAdapter := postgres_adapters.NewGormAdapter(db) // ? implement singleton deps container
	repo := infra.NewUserRepo(dbAdapter)
	service := app.NewUserService(repo)

	g.GET("user/:id", handle(getById, service))
	g.GET("user", handle(list, service))
	g.POST("user", handle(create, service))
	g.POST("user/:id/enable", handle(enable, service))
	g.POST("user/:id/disable", handle(disable, service))
}

func handle(m method, s ports.UserServicePort) gin.HandlerFunc { // TODO: implement auth service
	return func(c *gin.Context) {
		m(c, s)
	}
}

type method func(c *gin.Context, s ports.UserServicePort)
