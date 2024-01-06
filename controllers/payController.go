package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func CreateSessionController(ctx *gin.Context) {
	// Fetching the request body
	var requestBody struct {
		OrderID         *string
		Amount          int
		CustomerDetails types.JuspayCustomer
	}

	ctx.Bind(&requestBody)

}

func CheckOrderStatusController(ctx *gin.Context) {
	// Fetching the request body
	var requestBody struct {
		OrderID *string
	}

	ctx.Bind(&requestBody)
}
