package middlewares

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// CountPaidUsers counts the number of Paid Users
func CountPaidUsers() (int, error) {
	// Initialising the Context and Client
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Declaring the global variables
	iter := client.Collection("Users").Documents(ctx)
	var count int = 0

	// Looping through all the Users
	for {
		_, err := iter.Next()

		// Checking for the End of the list of Users
		if err == iterator.Done {
			break
		}
		if err != nil {
			return 0, err
		}
		// Incrementing the counter by 1
		count++
	}

	// Returning the final Count of all The Users
	return count, nil
}
