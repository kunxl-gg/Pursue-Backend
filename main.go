package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/routes"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// Initialising the necessary configurations
func init() {
	initialisers.LoadEnvVariables()
	initialisers.InitialiseFirebase()
}

func main() {
	// Initialising the routes
	r := gin.Default()

	// Grouping admin Routes together
	types.AdminRoutes = r.Group("/api/admin")
	types.UserRoutes = r.Group("/api/user")

	// Methods to call the Admin and User Routes
	routes.SetupAdminRoutes()
	routes.SetupUserRoutes()

	// Starting the server
	r.Run(":8080")
}
