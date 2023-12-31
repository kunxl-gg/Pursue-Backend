package routes

// Define all the Admin Routes here
import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	firebase_middleware "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/firebase"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func SetupAdminRoutes() {
	// Common middleware for all the Admin Routes
	types.AdminRoutes.Use(firebase_middleware.CheckIfAdmin)

	// Ping Route for Admin Routes
	types.AdminRoutes.GET("/", controllers.PingController)

	// Firebase (Users) call URLs
	types.AdminRoutes.GET("/userCount", controllers.CountUserController)
	types.AdminRoutes.GET("/paidUserCount", controllers.CountPaidUserController)
	types.AdminRoutes.GET("/chatbotSessionsCount", controllers.CountTotalChatbotSessionsController)
	types.AdminRoutes.GET("/users", controllers.ListOfUsersController)
	types.AdminRoutes.GET("/paidUsers", controllers.ListOfPaidUsersController)
	types.AdminRoutes.GET("/:userID", controllers.GetUserController)

	// Firebase (CareerDescription) call URLs
	types.AdminRoutes.POST("/addCareerDescription", controllers.AddCareerDescriptionToFirebaseController)
	types.AdminRoutes.GET("/careerDescription", controllers.GetCareerOptionsController)

	// URLs to update the Graph Database
	types.AdminRoutes.POST("/addNode", controllers.AddNodeController)
	types.AdminRoutes.POST("/updateNode", controllers.UpdateNodeController)

	// URLs to deal with the Admin User
	types.AdminRoutes.POST("/addAdminUser", nil)

	// URLs to addd routes to add question to the admin dashboard
	types.AdminRoutes.POST("/addToRepository", controllers.WriteToRepositoryController)
	types.AdminRoutes.GET("/readRepository/:id", controllers.ReadRepositoryController)

	// URLs to add question
	types.AdminRoutes.POST("/addQuestion", controllers.AddQuestionController)
	types.AdminRoutes.GET("/readQuestions", controllers.ReadQuestionController)
}
