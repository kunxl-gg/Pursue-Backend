package initialisers

import (
	"context"
	"fmt"
	"os"

	"github.com/mixpanel/mixpanel-go"
)

// InitialiseMixpanel initialises the Mixpanel API Client and a background context
func InitialiseMixpanel(event string, object string) error {

	projectToken := os.Getenv("MIXPANEL_PROJECT_TOKEN")
	fmt.Println(projectToken)

	// Initialising the Mixpanel API Client and a background context
	ctx := context.Background()
	client := mixpanel.NewApiClient(projectToken)

	// Adding an event to the Mixpanel dashboard
	err := client.Track(ctx, []*mixpanel.Event{
		client.NewEvent(event, "", map[string]interface{}{
			"Clicked Object": object,
		}),
	})

	// Checking if there is an error
	if err != nil {
		return err
	}

	return nil

}
