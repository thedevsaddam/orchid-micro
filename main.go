package main

import (
	_ "app/config"
	"app/routes"
	"os"
	"runtime"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	//set mode
	gin.SetMode(os.Getenv("APP_MODE")) //Caution: You must put the mode to "gin.ReleaseMode" when using in production

	//set the number of CPU processor will be used
	runtime.GOMAXPROCS(runtime.NumCPU())

	//instantiate route
	router := gin.Default()

	//initialize the routes
	routes.InitRoutes(*router)

	//route not found response
	router.NoRoute(func(c *gin.Context) {
		message := map[string]interface{}{"message": "The requested uri is not valid!", "code": 4004}
		c.JSON(404, gin.H{"status": 404, "error": message})
		return
	})

	//serve static files
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	//start the server
	router.Run(":" + os.Getenv("APP_PORT"))
}
