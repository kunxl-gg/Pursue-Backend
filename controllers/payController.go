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

// CreateCustomerController - Controller to Interact with the Juspay Server to Create Customer
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

// GetPaymentMethodsController - Controller Method to fetch the
func GetPaymentMethodsController(ctx *gin.Context) {
	resp, err := middlewares.GetPaymentMethods()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(
		http.StatusOK,
		resp,
	)
}

// CreateOrderController - Controller Method to Create an Order on the Juspay Server
func CreateOrderController(ctx *gin.Context) {
	var requestBody struct {
		Uuid     int
		Amount   string
		Currency string
	}

	// Reading Request Body
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Creating an Order
	resp, err := middlewares.CreateOrder(requestBody.Uuid, requestBody.Amount, requestBody.Currency)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(
		http.StatusInternalServerError,
		string(body))

}

// InitiatePaymentController - Controller Method to InitiatePayment Using API
func InitiatePaymentController() {

}
