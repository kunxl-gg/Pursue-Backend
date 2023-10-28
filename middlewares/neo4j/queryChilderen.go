package Neo4j

import (
	"context"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func QueryChildren(Node string) ([]string, error) {

	// Getting a background context
	ctx := context.Background()
	driver := initialisers.InitialiseNeo4jDB(ctx)
	defer driver.Close(ctx)

	// Making a query
	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (p:Person {name: $name}) -[:Recommends]-> (c:Person) RETURN c.name AS name",
		map[string]any{
			"name": Node,
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	// Checking for error in querying the DB
	if err != nil {
		return nil, err
	}

	var finalResponse []string
	for _, record := range result.Records {
		name, _ := record.Get("name")
		finalResponse = append(finalResponse, name.(string))
	}
	// Returning the final Result
	return finalResponse, nil

}
