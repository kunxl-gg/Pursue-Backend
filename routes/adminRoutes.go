package routes

// Define all the Admin Routes here
import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/controllers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func SetupRoutes() {
	types.AdminRoutes.GET("/", controllers.PingController)
}