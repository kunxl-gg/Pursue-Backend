package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	middlewares "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/payments"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"io"
	"net/http"
)

// This file Contains the Controllers for all Payment related Stuff

func GetCustomerController(ctx *gin.Context) {
	middlewares.GetCustomer()
}

// CreateCustomer - Controller to Interact with the Juspay Server to Create Customer
func CreateCustomerController(ctx *gin.Context) {

	// Fetching the context body
	var CustomerDetails types.JuspayCustomer
	err := ctx.Bind(&CustomerDetails)
	if err != nil {
		fmt.Println(err)
	}

	// Making a Fetch call to Juspay Server to create Customer
	resp, err := middlewares.CreateCustomer(*CustomerDetails.EmailAddress, *CustomerDetails.FirstName, *CustomerDetails.LastName, CustomerDetails.MobileNumber, CustomerDetails.CountryCode)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(
		http.StatusOK,
		body,
	)
}

func GetPaymentMethodsController(ctx *gin.Context) {

	// Getting All the Payment Methods
}
