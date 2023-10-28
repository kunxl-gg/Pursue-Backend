package controllers

import (
	"github.com/gin-gonic/gin"
	Neo4j "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/neo4j"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"net/http"
)

func QueryNodeController(ctx *gin.Context) {

	var NodeDetails types.Node
	ctx.Bind(&NodeDetails)

	result, _ := Neo4j.QueryChildren(*NodeDetails.NodeTitle)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"answer": result,
		},
	)
}
