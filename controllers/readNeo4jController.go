package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

func QueryNodeController(givenctx *gin.Context) {
	// Getting a background context
	ctx := context.Background()
	driver := initialisers.InitialiseNeo4jDB(ctx)
	defer driver.Close(ctx)

	// Making a query
	result, err := neo4j.ExecuteQuery(
		ctx,
		driver,
		"MATCH (p:Person {name: $name}) RETURN p.name AS name",
		map[string]any{
			"name": "Kunal Tiwari",
		},
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	if err != nil {
		panic(err)
	}

	// Iterating over the result
	var finalResponse string = " "
	for _, record := range result.Records {
		name, _ := record.Get("name")
		finalResponse += name.(string)
	}

	givenctx.JSON(
		http.StatusOK,
		gin.H{
			"answer": finalResponse,
		},
	)
}
