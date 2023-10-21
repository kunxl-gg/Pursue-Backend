package routes

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// Define all the User Routes Here

func SetupUserRoutes() {
	// URL to the Payment Gateway
	types.UserRoutes.GET("/pay", nil)

	// URL for querying information from the Graph DB
	types.UserRoutes.GET("/queryNode", controllers.QueryNodeController)

	// Calls to the Firebase DB from the Users End
	types.UserRoutes.POST("/updateStageForUser", nil)
	types.UserRoutes.POST("/addUser", nil)
}
