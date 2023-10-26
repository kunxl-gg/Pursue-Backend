package routes

// Define all the Admin Routes here
import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func SetupAdminRoutes() {
	// Common middleware for all the Admin Routes
	types.AdminRoutes.Use(middlewares.CheckIfAdmin)

	// URL to check if the server is up
	types.AdminRoutes.GET("/", controllers.PingController)

	// Firebase call URLs
	types.AdminRoutes.GET("/users", controllers.ListOfUsersController)
	types.AdminRoutes.GET("/paidUsers", controllers.ListOfPaidUsersController)
	types.AdminRoutes.GET("/:userID", nil)

	// URLs to update the Graph Database
	types.AdminRoutes.GET("/addNode", controllers.AddNodeController)
	types.AdminRoutes.GET("/updateNode", controllers.UpdateNodeController)

	// URLs to deal with the Admin User
	types.AdminRoutes.POST("/addAdminUser", nil)
}
