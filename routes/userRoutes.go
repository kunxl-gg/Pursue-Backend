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
	types.UserRoutes.POST("/createCustomer", controllers.CreateCustomerController)
	types.UserRoutes.GET("/getCustomer", nil)
	types.UserRoutes.GET("/getPaymentMethods", controllers.GetPaymentMethodsController)
	types.UserRoutes.POST("/createOrder", controllers.CreateOrderController)
	types.UserRoutes.GET("/initiateUPIPayment")

	// URL for querying information from the Graph DB
	types.UserRoutes.POST("/queryNode", controllers.QueryNodeController)

	// Calls to the Firebase DB from the Users End
	types.UserRoutes.POST("/addUser", controllers.AddUserController)
	types.UserRoutes.POST("/updateStageForUser/:userId", controllers.UpdateStageOfUserController)
	types.UserRoutes.POST("/updateSelectedOption/:userId", controllers.UpdateOptionsController)

	// Calls to Mixpanel
	types.UserRoutes.POST("/addEvent", controllers.TrackMixpanelController)
}
