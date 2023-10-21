package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"net/http"
)

// UpdateDataController Method to add Element to Firebase
func UpdateDataController(ctx *gin.Context) {

	// Global Variable for catching the FirebaseUser
	var requestBody types.FirebaseUser

	// Binding the incoming request to the request body
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "There was an Error in Reading the Request body ", err)
	}
	fmt.Println(requestBody.Name, requestBody.Stage, requestBody.IsPaidUser)

}
