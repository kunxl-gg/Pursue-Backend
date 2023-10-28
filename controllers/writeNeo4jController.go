package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	Neo4j "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/neo4j"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

// AddNodeController Method to add node to the Neo4j Database
func AddNodeController(ctx *gin.Context) {

	// Getting the response Body
	var NodeDetails types.Node
	err := ctx.Bind(&NodeDetails)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Adding Node to the DB
	Neo4j.AddNode(*NodeDetails.NodeTitle)
}

// UpdateNodeController Method to update the children of a node in the Neo4j DB
func UpdateNodeController(givenctx *gin.Context) {

	// Getting background context
	ctx := context.Background()
	driver := initialisers.InitialiseNeo4jDB(ctx)
	defer driver.Close(ctx)

	// Making an update
	_, err := neo4j.ExecuteQuery(
		ctx,
		driver,
		"MATCH (p:Person {name: $name}) MERGE (p) -[:LIKES]->(:Person {name:$tech})",
		map[string]any{
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
