package middlewares 

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func CreateOrder(uuid int, amount string, currency string) (*http.Response, error) {

	client := &http.Client{}

	// Define the URL and payload
	urlStr := "https://api.juspay.in/orders"
	data := url.Values{}
	data.Set("order_id", strconv.Itoa(uuid))
	data.Set("amount", amount)
	data.Set("currency", currency)

	// Create a new POST request
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("version", "2023-06-30")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-merchantid", os.Getenv("MERCHANT_ID"))
	req.Header.Set("Authorization", "Basic NjQ0Nzg4RTlDNDU0MDVEOTc4NDg1NDk1QjA3RTVD")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, err
	}

	// Returning the Response
	return resp, nil
}
