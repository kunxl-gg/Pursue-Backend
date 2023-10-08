package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/helpers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func AddStandAloneNode(givenctx *gin.Context) {
	// Getting the background Context
	ctx := context.Background()
	driver := helpers.MakeConnectionWithDb(ctx)
	defer driver.Close(ctx)


}

func AddNodeController(givenctx *gin.Context) {
	var requestBody struct {
		name string
	}

	givenctx.Bind(&requestBody)
	ctx := context.Background()
	driver := helpers.MakeConnectionWithDb(ctx)
	defer driver.Close(ctx)

	_, err := neo4j.ExecuteQuery(ctx, driver, "CREATE (p:Person {name: $name} ) -[:LIKES]->(:Person {name:$simp}) RETURN p", map[string]any{
		"simp": "Kunal Tiwari",
		"name": "Jesika Rai",
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
		"MATCH (p:Person {name: $name}) MERGE (p) -[:LIKES]->(:Person {name:$simp})",
		map[string]any {
			"simp": "Dhruv Dighe",
			"name": "Siya Bhadra",
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
