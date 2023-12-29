package controllers

import (
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
	id, err := chatbot.AddStudentQuestion(*requestBody.Section, *requestBody.Questions, requestBody.Options)

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
