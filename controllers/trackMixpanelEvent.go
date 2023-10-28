package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"net/http"
)

// TrackMixpanelController - Method to interact with Mixpanel
func TrackMixpanelController(ctx *gin.Context) {

	// Capturing the request body
	var MixpanelEvent struct {
		Event  *string
		Object *string
	}
	err := ctx.Bind(&MixpanelEvent)
	fmt.Println(MixpanelEvent.Event, MixpanelEvent.Object)

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Tracking an event using Mixpanel
	err = initialisers.InitialiseMixpanel(*MixpanelEvent.Event, *MixpanelEvent.Object)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Sending a 200 response if everything goes right
	ctx.String(http.StatusOK, "Successfully Added Event")

}
