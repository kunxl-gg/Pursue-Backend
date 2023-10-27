package initialisers

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"os"
)

// InitialiseNeo4jDB - Method to make connection with neo4j database
func InitialiseNeo4jDB(ctx context.Context) neo4j.DriverWithContext {

	// DB Credentials
	dbUri := os.Getenv("NEO4JDB_URL")
	dbUser := os.Getenv("NEO4JDB_USERNAME")
	dbPassword := os.Getenv("NEO4JDB_PASSWORD")

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
