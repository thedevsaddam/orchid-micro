package routes

import (
	"github.com/gin-gonic/gin"

	"app/controllers/users"
	"app/controllers/welcome"
	"app/middlewares/cors"
	"app/middlewares/example"
)

func InitRoutes(router gin.Engine) {
	//use middleware
	router.Use(cors.CORSMiddleware())

	//root route
	router.GET("/", welcome.Welcome)
	router.GET("/users", users.FetchUsers)

	//route group
	v1 := router.Group("/api/v1").Use(example.Example())
	{
		v1.GET("/", welcome.Hello)
	}
}
