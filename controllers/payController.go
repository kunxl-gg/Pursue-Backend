package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares"
)

func PayController(ctx *gin.Context) {
	middlewares.CreateCustomer()
}