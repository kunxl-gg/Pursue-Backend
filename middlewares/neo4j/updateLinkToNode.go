package Neo4j

import (
	"context"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func UpdateLinkToNode(parent string, child string) error {

	// Getting background context
	ctx := context.Background()
	driver := initialisers.InitialiseNeo4jDB(ctx)
	defer driver.Close(ctx)

	// Making an update
	_, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (p:Person {name: $parent}) MERGE (p) -[:Recommends]->(:Person {name: $child})",
		map[string]any{
			"child":  child,
			"parent": parent,
		}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	return nil
}
