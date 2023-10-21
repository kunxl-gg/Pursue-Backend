package controllers

// This File has the Controllers for all the Read Operations to Firebase

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CountUserController Function to fetch Total User Count
func CountUserController(ctx *gin.Context) {
	count := CountTotalUsers()
	ctx.String(http.StatusOK, strconv.Itoa(count))
}

// CountPaidUserController Method to fetch the total Count of Paid Users
func CountPaidUserController(ctx *gin.Context) {

}

// ListOfUsersController Method to fetch Total User  List
func ListOfUsersController(ctx *gin.Context) {
	users := ListOfUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		})
}

// ListOfPaidUsersController Method to fetch the list of Paid Users
func ListOfPaidUsersController(ctx *gin.Context) {
	users := ListOfPaidUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		})
}
