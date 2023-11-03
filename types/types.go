package types

import (
	"github.com/gin-gonic/gin"
	"os"
)

var AdminRoutes *gin.RouterGroup
var UserRoutes *gin.RouterGroup

type FirebaseUser struct {
	Name       *string
	IsPaidUser bool
	Stage      int
}

type Node struct {
	NodeTitle *string
}

type JuspayCustomer struct {
	EmailAddress *string
	FirstName    *string
	LastName     *string
	MobileNumber int
	CountryCode  int
}

var MerchantID string = os.Getenv("MERCHANT_ID")
var JuspayAPIKey string = os.Getenv("JUSPAY_API_KEY")
var Base64encodedAPIKey = os.Getenv("Base64EncodedAPIKey")
