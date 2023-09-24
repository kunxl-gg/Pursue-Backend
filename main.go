package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
}

func main() {
	// Initialising the routes
	r := gin.Default()

	// Defining all the routes
	r.GET("/", nil)
	r.GET("/pay", nil)
	r.GET("/authenticate", nil)

	// Starting the server
	r.Run(":8080")
}
