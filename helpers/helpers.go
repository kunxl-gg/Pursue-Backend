package helpers

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Helper Function to Check Error 
func CheckError(err error) {
	panic(err)
}

// Helper Function to Check Error with Custom Message
func CheckErrorWithCustomMessage(err error) {
	panic(err)
}

// Helper Function to make connection with database
func MakeConnectionWithDb(ctx context.Context) neo4j.DriverWithContext {
	// DB Credentials
	dbUri := "neo4j://localhost"
	dbUser := "neo4j"
	dbPassword := "secretgraph"

	// Creating a Driver for the Database
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		panic(err)
	}

	// Verifying the Connection
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	// returning the driver
	return driver
} 