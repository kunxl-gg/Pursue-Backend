package initialisers

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// InitialiseFirebase initialises the Firebase API Client and a background context
func InitialiseFirebase() (context.Context, *firestore.Client) {

	// Define the background Context and Configuration
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: "amritcounsellor-17082",
	}
	
	// Getting the path of the current working directory
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Configuring the secret key
	opt := option.WithCredentialsFile(curDir + "/secret-key.json")

	// Initialise the Ap
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatal(err)
	}

	// Initialise the client to interact with the firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Return both the values
	return ctx, client
}
