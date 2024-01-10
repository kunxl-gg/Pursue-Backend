package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/chatbot"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func AddQuestionController(ctx *gin.Context) {

	// Reading the request body
	var requestBody types.ChatBotQuestion
	err := ctx.Bind(&requestBody)
	if err != nil {
		log.Fatal(err)
	}

	// Making a request to the middleware
	id, err := chatbot.AddQuestion(*requestBody.ID, *requestBody.Section, *requestBody.Questions, requestBody.Options)

	ctx.String(http.StatusOK, "Added ChatBotQuestion to the Database", id)
}

func ReadQuestionController(ctx *gin.Context) {
	// Call the middleware function
	data, err := chatbot.ReadQuestion()
	if err != nil {
		log.Fatal(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Return a response
	ctx.JSON(
		http.StatusOK,
		data,
	)
}

// Method to Edit Question
func EditQuestionController(ctx *gin.Context) {
	// Binding the requestBody to the variable
	var requestBody types.ChatBotQuestion
	ctx.Bind(&requestBody)
	fmt.Println(*requestBody.ID)

	// Pass the requestBody into some sort of middleware
	resp, err := chatbot.EditQuestion(*requestBody.ID, *requestBody.Section, *requestBody.Questions, requestBody.Options)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Return a Response
	ctx.String(http.StatusOK, resp)
}

// Method to Delete Question
func DeleteQuestionController(ctx *gin.Context) {
	// Binding the requestBody to the variable
	var requestBody types.ChatBotQuestion
	ctx.Bind(&requestBody)
	fmt.Println(requestBody.Section)

	// Pass the requestBody into some sort of middleware
	resp, err := chatbot.DeleteQuestion(*requestBody.ID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	// Return a Response

	ctx.String(http.StatusOK, resp)
}
