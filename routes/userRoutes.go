package routes

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// Define all the User Routes Here

func SetupUserRoutes() {

	// Ping Route for User Routes
	types.UserRoutes.GET("/", controllers.PingController)

	// URL to the Payments System
	types.UserRoutes.POST("/createSession", controllers.CreateSessionController)
	types.UserRoutes.POST("/checkOrderStatus", controllers.CheckOrderStatusController)

	// Calls to the Firebase DB from the Users End
	types.UserRoutes.POST("/addUser", controllers.AddUserController)
	types.UserRoutes.POST("/updateStageForUser/:userId", controllers.UpdateStageOfUserController)
	types.UserRoutes.POST("/updateSelectedOption/:userId", controllers.UpdateOptionsController)
	types.UserRoutes.POST("/getCareerOption", controllers.GetCareerOptionController)

	// Calls to Mixpanel
	types.UserRoutes.POST("/addEvent", controllers.TrackMixpanelController)
}
