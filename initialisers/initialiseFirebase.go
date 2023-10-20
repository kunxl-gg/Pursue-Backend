package initialisers

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

func InitialiseFirebase() (context.Context, *firestore.Client) {

	// Define the background Context and Configuration
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: "amritcounsellor-17082",
	}

	// Configuring the secret key
	opt := option.WithCredentialsFile("/Users/kunaltiwari/Desktop/Projects/Amrit-Career-Counsellor/secret-key.json")

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
