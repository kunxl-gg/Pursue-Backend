package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

// Initialising the necessary configurations
func init() {
	initialisers.LoadDB()
	initialisers.LoadEnvVariables()
}

func main() {
	// Initialising the routes
	r := gin.Default()

	// Defining all the routes
	r.GET("/", controllers.PingController)
	r.GET("/pay", nil)
	r.GET(`/authenticate`, nil)
	r.GET("/queryNode", controllers.QueryNodeController)
	r.POST("/addNode", controllers.AddNodeController)
	r.POST("/updateNode", controllers.UpdateNodeController)

	// Starting the server
	r.Run(":8080")
}
