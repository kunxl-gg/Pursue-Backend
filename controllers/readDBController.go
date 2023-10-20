package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
	"strconv"
)

func CountTotalUsers() int {
	ctx, client := initialisers.InitialiseFirebase()

	// Getting the itreator for the db
	iter := client.Collection("Users").Documents(ctx)

	// Looping through the iterator
	var count int = 0
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			return count
		}
		if err != nil {
			log.Fatal(err)
		}
		count++
	}
}

// Make a function to count the total Number of Paid Users
func ListOfUsers() {
	ctx, client := initialisers.InitialiseFirebase()

	iter := client.Collection("Users").Documents(ctx)

	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data.Data()["First Name"])
	}
}

// Make a function to Fetch the Data of a particular Person
func ListOfPaidUsers() []map[string]interface{} {
	ctx, client := initialisers.InitialiseFirebase()

	iter := client.Collection("Users").Documents(ctx)

	var Users []map[string]interface{}
	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if data.Data()["Paid"] == "false" {
			Users = append(Users, data.Data())
		}
	}

	return Users
}

// CountUserController Make a function to fetch any data from the given DB
func CountUserController(ctx *gin.Context) {
	var count int = CountTotalUsers()
	ctx.String(http.StatusOK, strconv.Itoa(count))
}

// ListOfUsersController Controller for the list of Users
func ListOfUsersController(ctx *gin.Context) {
	ListOfUsers()
	ctx.String(http.StatusOK, "This is working")
}

// ListOfPaidUsersController Controller for the list of Paid Users
func ListOfPaidUsersController(ctx *gin.Context) {
	users := ListOfPaidUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		},
	)
}
