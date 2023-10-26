package middlewares

import "os"

func GetPaymentMethods() {
	
	// JustPay API URL and Merchant ID
	apiURL := os.Getenv("JUSPAY_SERVER_URL")
	merchantID := os.Getenv("MERCHANT_ID")
	
}