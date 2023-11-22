package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"log"
	"net/http"
)

// AddUserController - Method to add User to Firestore
func AddUserController(ctx *gin.Context) {

	// Initialising a Global variable for the Firebase User
	var user types.FirebaseUser

	// Binding the incoming JSON to the User variable
	err := ctx.Bind(&user)
	if err != nil {
		log.Fatal("There was an error in reading the Incoming Json", err)
	}

	// Adding User to Firestore
	userId, err := middlewares.AddUserToFirebase(*user.Name, user.IsPaidUser, user.Stage, user.Options, user.FinalCareerOptions)
	if err != nil {
		log.Fatal("There was an error in adding user to firestore", err)
	}

	// Returning the final userID
	ctx.String(http.StatusOK, string(userId))

}

// DeleteUserController - Method to delete User from Firestore
func DeleteUserController(ctx *gin.Context) {

	// Fetching the UserId from the URL path
	userId := ctx.Param("userId")

	// Deleting the User from Firestore
	err := middlewares.DeleteUser(userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Deleted User ", userId)

}

// UpdateStageOfUserController - Method to update value of stage in Firestore
func UpdateStageOfUserController(ctx *gin.Context) {

	// Fetching the UserId from the URL path
	var userId string = ctx.Param("userId")

	var stage struct {
		Stage int
	}
	ctx.Bind(&stage)
	fmt.Println(stage)

	// Updating Stage for User
	err := middlewares.UpdateStageOfUser(stage.Stage, userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Updated Stage for User ", userId)

}

func UpdateOptionsController(ctx *gin.Context) {
	
	// Fetching the UserId from the URL path
	var userId string = ctx.Param("userId")
	fmt.Println(userId)

	var options struct {
		NewOptionSelected *string
	}
	ctx.Bind(&options)

	// Updating Stage for User
	err := middlewares.UpdateSelectedOption(userId, *options.NewOptionSelected)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Updated Options for User ", userId)
}
