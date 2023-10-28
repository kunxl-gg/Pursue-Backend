package middlewares

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

// AddUserToFirebase Method to add a User to Firestore
func AddUserToFirebase(name string, isPaidUser bool, stage int) (string, error) {

	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding data to the DB
	doc, _, err := client.Collection("Users").Add(ctx, map[string]interface{}{
		"First Name": name,
		"Stage":      stage,
		"IsPaidUser": isPaidUser,
	})
	if err != nil {
		return "", err
	}

	return doc.ID, nil

}
