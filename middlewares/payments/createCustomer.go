package middlewares 

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func CreateCustomer(emailID string, firstName string, secondName string, mobileNumber int, countryCode int) (*http.Response, error) {
	// Fetching the credentials for the API request
	apiURL := os.Getenv("JUSPAY_SERVER_URL") + "/customers/"
	apiKey := os.Getenv("JUSPAY_API_KEY")
	merchantID := os.Getenv("MERCHANT_ID")

	// Fetching the value of request Body

	// Setting values for metadata
	data := url.Values{}
	data.Set("object_reference_id", emailID)
	data.Set("mobile_number", strconv.Itoa(mobileNumber))
	data.Set("email_address", emailID)
	data.Set("first_name", firstName)
	data.Set("last_name", secondName)
	data.Set("mobile_country_code", strconv.Itoa(countryCode))
	data.Set("options.get_client_auth_token", "true")

	// Initiating the Request
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Setting headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-merchantid", merchantID)
	req.SetBasicAuth(apiKey, "")

	// Making the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	return resp, nil

}
