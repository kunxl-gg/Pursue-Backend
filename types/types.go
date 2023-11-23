package types

import (
	"os"

	"github.com/gin-gonic/gin"
)

var AdminRoutes *gin.RouterGroup
var UserRoutes *gin.RouterGroup

type FirebaseUser struct {
	Name               *string
	IsPaidUser         bool
	Stage              int
	Options            []string
	FinalCareerOptions []string
}

type FirebaseAdminUser struct {
	Name  *string
	Email *string
}

type FirebaseCareerOption struct {
	Name            *string
	Description     *string
	CareerPathSteps []string
	AverageSalaries []map[string][]int
	TopColleges     []TopCollge
}

type TopCollge struct {
	Name  *string
	Image *string
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
