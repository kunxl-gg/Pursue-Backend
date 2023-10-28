package Neo4j

import (
	"context"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

// AddNode - Method to add Node to database
func AddNode(NodeTitle string) {

	// Getting the Background Context
	ctx := context.Background()

	// Making connection with Neo4j
	driver := initialisers.InitialiseNeo4jDB(ctx)
	defer driver.Close(ctx)

	// Adding a Node in the Neo4j Database
	_, err := neo4j.ExecuteQuery(ctx, driver, "CREATE (p:Person {name: $name}) RETURN p", map[string]any{
		"name": NodeTitle,
	}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		log.Fatal(err)
	}

	// After everything goes fine
	return

}
