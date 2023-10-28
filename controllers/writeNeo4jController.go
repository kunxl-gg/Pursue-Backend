package controllers

import (
	"github.com/gin-gonic/gin"
	Neo4j "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/neo4j"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
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
	err = Neo4j.AddNode(*NodeDetails.NodeTitle)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending a 200 Response if everything goes fine
	ctx.String(http.StatusOK, "Added a Node ", NodeDetails.NodeTitle)

}

// UpdateNodeController Method to update the children of a node in the Neo4j DB
func UpdateNodeController(ctx *gin.Context) {

	// Fetching the request Body
	var link struct {
		Parent string
		Child  string
	}
	ctx.Bind(&link)

	// Updating the link
	_ = Neo4j.UpdateLinkToNode(link.Parent, link.Child)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"answer": "Updated",
		},
	)

}
