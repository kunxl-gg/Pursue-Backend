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

	// Ping Route for Admin Routes
	types.AdminRoutes.GET("/", controllers.PingController)

	// Firebase (Users) call URLs
	types.AdminRoutes.GET("/userCount", controllers.CountUserController)
	types.AdminRoutes.GET("/paidUserCount", controllers.CountPaidUserController)
	types.AdminRoutes.GET("/users", controllers.ListOfUsersController)
	types.AdminRoutes.GET("/paidUsers", controllers.ListOfPaidUsersController)
	types.AdminRoutes.GET("/:userID", controllers.GetUserController)
	//types.AdminRoutes.GET("/:userId", controllers.DeleteUserController)

	// Firebase (CareerDescription) call URLs
	types.AdminRoutes.POST("/addCareerDescription", controllers.AddCareerDescriptionToFirebaseController)
	types.AdminRoutes.GET("/careerDescription", controllers.GetCareerOptionsController)

	// URLs to update the Graph Database
	types.AdminRoutes.POST("/addNode", controllers.AddNodeController)
	types.AdminRoutes.POST("/updateNode", controllers.UpdateNodeController)

	// URLs to deal with the Admin User
	types.AdminRoutes.POST("/addAdminUser", nil)
}
