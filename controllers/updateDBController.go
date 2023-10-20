package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"net/http"
)

func UpdateDataToDB(name string, isPaidUser string, stage string) error {
	ctx, client := initialisers.InitialiseFirebase()

	// Adding data to the DB
	_, _, err := client.Collection("Users").Add(ctx, map[string]interface{}{
		"First Name": name,
		"Stage":      stage,
		"Paid":       isPaidUser,
	})

	if err != nil {
		return err
	}

	return nil
}

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
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err,
			})
	}
}
