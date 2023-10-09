package routes

// Define all the Admin Routes here
import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func SetupRoutes() {
	types.AdminRoutes.Use(middlewares.CheckIfAdmin)
	types.AdminRoutes.GET("/", controllers.PingController)
}