package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type JuspayResponse struct {
	OrderID     string
	RedirectURL string
}

func PayController(context *gin.Context) (*JuspayResponse, error) {
	var orderBody struct {
		OrderID string
		Amount  string
	}

	context.Bind(&orderBody)

	const (
		juspayBaseURL = "https://sandbox.juspay.in" // Use sandbox for testing
		merchantID    = "YOUR_MERCHANT_ID"          // Replace with your merchant ID
		apiKey        = "YOUR_API_KEY"              // Replace with your API Key
	)

	url := juspayBaseURL + "/orders"

	payload := strings.NewReader(fmt.Sprintf("order_id=%s&amount=%.2f", orderBody.OrderID, orderBody.Amount))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+apiKey) // Typically you would use a Base64 encoded API key

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response JuspayResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
