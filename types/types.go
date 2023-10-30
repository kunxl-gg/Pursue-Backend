package types

import (
	"github.com/gin-gonic/gin"
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
