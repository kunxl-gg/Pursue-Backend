package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/helpers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func QueryNodeController(givenctx *gin.Context) {
	// Getting a background context
	ctx := context.Background()
	driver := helpers.MakeConnectionWithDb(ctx)
	defer driver.Close(ctx)

	// Making a query
	result, err := neo4j.ExecuteQuery(
		ctx, 
		driver, 
		"MATCH (p:Person {name: $name}) RETURN p.name AS name",
		map[string]any {
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
		finalResponse  += name.(string)
	}

	givenctx.JSON(
		http.StatusOK,
		gin.H{
			"answer": finalResponse,
		},
	)
}

func AddNodeController(givenctx *gin.Context) {
	var requestBody struct {
		name string
	}

	givenctx.Bind(&requestBody)
	ctx := context.Background()
	driver := helpers.MakeConnectionWithDb(ctx)
	defer driver.Close(ctx)

	_, err := neo4j.ExecuteQuery(ctx, driver, "CREATE (p:Person {name: $name}) -[:LIKES]->(:Person {name:$tech}) RETURN p", map[string]any{
		"tech": "Node.js",
		"name": "Kunal Tiwari",
	}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	helpers.CheckError(err)

	givenctx.JSON(
		http.StatusOK,
		gin.H{
			"answer":          requestBody.name,
		},
	)
}

func UpdateNodeController(givenctx *gin.Context) {
	
	// Getting background context
	ctx := context.Background()
	driver := helpers.MakeConnectionWithDb(ctx)
	defer driver.Close(ctx)

	// Making an update
	_, err := neo4j.ExecuteQuery(
		ctx, 
		driver, 
		"MATCH (p:Person {name: $name}) MERGE (p) -[:LIKES]->(:Person {name:$tech})",
		map[string]any {
			"tech": "Flutter",
			"name": "Kunal Tiwari",
		}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil { 
		panic(err)
	}

	givenctx.JSON( 
		http.StatusOK,
		gin.H{
			"answer": "Updated",
		},
	)

}
