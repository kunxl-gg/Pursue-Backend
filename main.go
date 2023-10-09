package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/routes"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// Initialising the necessary configurations
func init() {
	initialisers.LoadDB()
	initialisers.LoadEnvVariables()
}

func main() {
	// Initialising the routes
	r := gin.Default()

	// Grouping admin Routes together
	types.AdminRoutes = r.Group("/api/admin")
	types.UserRoutes = r.Group("/api/user")

	// Calling the Function SetupRoutes
	routes.SetupRoutes()

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
