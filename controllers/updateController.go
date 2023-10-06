package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/helpers"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

func UpdateController(givenctx *gin.Context) {
	var requestBody struct {
		name string
	}

	givenctx.Bind(&requestBody)
	ctx := context.Background()
	dbUri := "neo4j://localhost"
	dbUser := "neo4j"
	dbPassword := "secretgraph"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		panic(err)
	}
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	result, err := neo4j.ExecuteQuery(ctx, driver, "MERGE (p:Person {name: $name}) RETURN p", map[string]any{
		"name": "Kunal Tiwari",
	}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))

	helpers.CheckError(err)

	givenctx.JSON(
		http.StatusOK,
		gin.H{
			"message":         result.Summary.Counters().NodesCreated(),
			"followUpMessage": result.Summary.ResultAvailableAfter(),
			"answer":          requestBody.name,
		},
	)
}
