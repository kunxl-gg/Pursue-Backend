package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/chatbot"
	"net/http"
)

func WriteToRepository(ctx *gin.Context) {
	// Reading the request body
	var requestBody struct {
		Questions []string
		Answers   []string
	}

	err := ctx.Bind(&requestBody)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	// Call the function
	id, err := chatbot.AddQuestionInRepository(requestBody.Questions, requestBody.Answers)
	fmt.Println(id, err)

	// Returning the final response Id of the object created
	ctx.String(http.StatusOK, id)
}

