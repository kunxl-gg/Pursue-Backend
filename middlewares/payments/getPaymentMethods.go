package middlewares

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetPaymentMethods - Gets the list of Payment Methods
func GetPaymentMethods() (string, error) {

	// JustPay API URL and Merchant ID
	juspayServerURL := os.Getenv("JUSPAY_SERVER_URL")
	merchantID := os.Getenv("MERCHANT_ID")

	// Final URL for fetching Payment Methods
	url := juspayServerURL + "/merchants/" + merchantID + "/paymentmethods"
	fmt.Println(url)

	// Creating a New Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	// Adding MerchantID to request Header
	req.Header.Add("x-merchantid", merchantID)
	req.Header.Add("Content-Type", "application/json")

	// Creating a Client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Closing the Response Body
	defer resp.Body.Close()

	// Reading the Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(string(body))

	// Printing the Response
	return string(body), nil
}
