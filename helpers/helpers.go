package helpers

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CheckError(err error) {
	fmt.Println(err)
	panic(err)
}

func CheckErrorWithCustomMessage(err error) {
	panic(err)
}

// Helper Function to make connection with database
func MakeConnectionWithDb(ctx context.Context) neo4j.DriverWithContext {

	dbUri := "neo4j://localhost"
	dbUser := "neo4j"
	dbPassword := "secretgraph"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		panic(err)
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	// returning the driver
	return driver
} 