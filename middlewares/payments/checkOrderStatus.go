package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func CreateOrder(orderId string) {
	// Creating an HTTP client
	client := http.Client{}

	orderUrl := os.Getenv("JUSPAY_SERVER_URL") + "/order/" + orderId
	merchantId := os.Getenv("MERCHANT_ID")
	currentDate := time.Now().Format("2024-01-30")

	// Setting headers to the request
	req, err := http.NewRequest("POST", orderUrl, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Setting the values of Headers
	req.Header.Set("version", currentDate)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-merchantid", merchantId)

	// Setting up Basic Authentication
	req.SetBasicAuth("", "")

	// Making the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
