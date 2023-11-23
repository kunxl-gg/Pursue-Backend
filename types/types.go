package types

import (
	"os"
	"github.com/gin-gonic/gin"
)

// Types for the Admin Routes and User Routes
var AdminRoutes *gin.RouterGroup
var UserRoutes *gin.RouterGroup

// Types for Firebase Schema
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
	Skills          []Skills
	Courses         []Courses
}

type Courses struct {
	Title       *string
	Description *string
	Image       *string
}

type Skills struct {
	Title       *string
	Description *string
}

type TopCollge struct {
	Name  *string
	Image *string
}

// Types for Neo4j DB
type Node struct {
	NodeTitle *string
}

// Types for the Payment Gateway
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
