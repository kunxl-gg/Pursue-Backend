package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"io"
	"net/http"
	"strconv"
)

// CreateSession: Middleware function to create a payment session
func CreateSession(orderId string, amount int, customerCredentials types.JuspayCustomer, description string) error {
	// Create a new Http client
	client := &http.Client{}

	// Defining our Request parameters
	requestUrl := "https://api.juspay.in"
	requestData := map[string]interface{}{
		"order_id":               orderId,
		"amount":                 strconv.Itoa(amount),
		"customer_id":            *customerCredentials.CustomerId,
		"customer_email":         *customerCredentials.EmailAddress,
		"customer_phone":         strconv.Itoa(customerCredentials.MobileNumber),
		"payment_page_client_id": "pursue",
		"action":                 "paymentPage",
	}

	// Serialising the above request body to json
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return err
	}

	// Maing a post request programmatically
	resp, err := client.Post(requestUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Reading the request body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(body)
	return nil
}
