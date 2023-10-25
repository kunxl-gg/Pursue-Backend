package initialisers

import (
	"context"

	"github.com/mixpanel/mixpanel-go"
)

// InitialiseMixpanel initialises the Mixpanel API Client and a background context
func InitialiseMixpanel()  (*mixpanel.ApiClient, context.Context) {
	
	// Initialising the Mixpanel API Client and a background context 
	ctx := context.Background()
	mp := mixpanel.NewApiClient("")
	
	// Returning the context and mixpanel API Client
	return mp, ctx

}