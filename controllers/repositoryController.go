package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/chatbot"
	"log"
	"net/http"
)

// ReadRepositoryController: Controller to fetch all the repository Details
func ReadRepositoryController(ctx *gin.Context) {
	// Reading the parameter from the URl
	id := ctx.Param("id")
	fmt.Println(id)

	// Making a call to the backend function
	data, err := chatbot.ReadRepository(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Returning the final response as a string
	ctx.JSON(
		http.StatusOK,
		data,
	)
}

// WriteToRepository: Controller to feed in repository details to the database
func WriteToRepositoryController(ctx *gin.Context) {
	// Reading the request body
	var requestBody struct {
		DatabaseTable     *string
		Parameters        []string
		CareerSuggestions []string
	}

	err := ctx.Bind(&requestBody)
	if err != nil {
		log.Fatal(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Call the function
	id, err := chatbot.AddUserChoicesInRepository(*requestBody.DatabaseTable, requestBody.Parameters, requestBody.CareerSuggestions)
	fmt.Println(id, err)

	// Returning the final response Id of the object created
	ctx.String(http.StatusOK, id)
}


