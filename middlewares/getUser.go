package middlewares

import (
	"fmt"

	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

// MIDDELWARE: GetUser - Get the User from Firestore using the userID
func GetUser(userID string) {
	// Client and context 
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Get the User from Firestore
	doc, err := client.Collection("Users").Doc(userID).Get(ctx)
	if err != nil {
		fmt.Println(err)
	}

	// Get the Data from the Document
	data := doc.Data()
	fmt.Println(data)

}