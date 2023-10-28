package middlewares

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// MIDDLWARE: GetCustomer - Check if the we have an existing customer witht he given customerID
func GetCustomer() {

	// JuspayServer URL, API KEY and MerchantID
	juspayAPIURL := os.Getenv("JUSPAY_SERVER_URL")
	juspayAPIKey := os.Getenv("JUSPAY_API_KEY")
	merchantID := os.Getenv("MERCHANT_ID")
	fmt.Println(juspayAPIURL, juspayAPIKey)

	// Final URL 
	url := juspayAPIURL + "/customers/"

	// Creating a New Request 
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return 
	}

	// Adding MerchantID to request Header 
	req.Header.Add("x-merchantid", merchantID)
	req.Header.Add("Content-Type", "application/json")

	// Adding Basic Auth using the API Key
	req.SetBasicAuth(juspayAPIKey, "")

	// Creating a Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 
	}

	// Closing the Response Body
	defer resp.Body.Close()

	// Reading the Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return 
	}

	// Printing the Response
	fmt.Println(string(body))

}
