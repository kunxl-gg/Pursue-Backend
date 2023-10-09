package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckIfAdmin(ctx *gin.Context) { 
	// Checking if the user is admin or not
	var requestBody struct { 
		IsAdmin bool `json:"isAdmin"`
	}

	ctx.Bind(&requestBody)

	if requestBody.IsAdmin {
		ctx.Next()
	}else {
		// In case the user is not admin then abort the request
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}