package middlewares

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func CreateCustomer() {
	apiURL := os.Getenv("JUSPAY_SERVER_URL") + "/customers/"
	apiKey := os.Getenv("JUSPAY_API_KEY")
	merchantID := os.Getenv("MERCHANT_ID")

	data := url.Values{}
	data.Set("object_reference_id", "customer@gmail.com")
	data.Set("mobile_number", "9988776655")
	data.Set("email_address", "customer@gmail.com")
	data.Set("first_name", "John")
	data.Set("last_name", "Smith")
	data.Set("mobile_country_code", "91")
	data.Set("options.get_client_auth_token", "true")

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-merchantid", merchantID)
	req.SetBasicAuth(apiKey, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

}
