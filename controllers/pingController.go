package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pursuse Says Hello!")
}
