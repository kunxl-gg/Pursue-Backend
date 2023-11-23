package controllers

// This File has the Controllers for all the Read Operations to Firebase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares"
	firebase_middleware "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/firebase"
)

// CountUserController Function to fetch Total User Count
func CountUserController(ctx *gin.Context) {
	count := middlewares.CountTotalUsers()
	ctx.String(http.StatusOK, strconv.Itoa(count))
}

// CountPaidUserController Method to fetch the total Count of Paid Users
func CountPaidUserController(ctx *gin.Context) {
	count, err := middlewares.CountPaidUsers()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.String(http.StatusOK, strconv.Itoa(count))
}

// ListOfUsersController Method to fetch Total User  List
func ListOfUsersController(ctx *gin.Context) {
	users := middlewares.ListOfUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		})
}

// ListOfPaidUsersController Method to fetch the list of Paid Users
func ListOfPaidUsersController(ctx *gin.Context) {
	users := middlewares.ListOfPaidUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		})
}

// GetUserController Method to fetch a single User
func GetUserController(ctx *gin.Context) {

	// Fetching the UserId from the path parameter
	userId := ctx.Param("userID")

	// Getting the user using the userID
	doc, err := middlewares.GetUserFromFirebase(userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Returning Final Json when everything goes fine
	ctx.JSON(
		http.StatusOK,
		doc.Data(),
	)
}

// GetCareerOptionsController Method to fetch the Career Options of a User
func GetCareerOptionsController(ctx *gin.Context) {

	description := firebase_middleware.ReadCareerDescriptionFromFirebase()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": description,
		},	
	)
}