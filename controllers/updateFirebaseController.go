package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func UpdateDataController(ctx *gin.Context) {

	var requestBody struct {
		name       string
		isPaidUser string
		stage      string
	}

	ctx.Bind(&requestBody)

	fmt.Println(requestBody.name, requestBody.stage, requestBody.isPaidUser)

	err := UpdateDataToDB(requestBody.name, requestBody.isPaidUser, requestBody.stage)
	if err != nil {
		fmt.Println(err)
	}
}
