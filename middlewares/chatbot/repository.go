package chatbot

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// ReadRepository: reads all the entries from the given repository
func ReadRepository(DatabaseTitle string) ([]map[string]interface{}, error) {
	// Initialising ctx and client for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Making an iterator for the given Table
	iter := client.Collection(DatabaseTitle).Documents(ctx)
	repositoryEntries := make([]map[string]interface{}, 0)
	for {
		data, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		repositoryEntries = append(repositoryEntries, data.Data())
	}

	// Returning all the repository entries
	return repositoryEntries, nil

}

func AddUserChoicesInRepository(DatabaseTitle string, Parameters []string, CareerOptions []string) (string, error) {
	// Initialising ctx and client for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding the Entry to Firebase
	doc, _, err := client.Collection(DatabaseTitle).Add(ctx, map[string]interface{}{
		"Parameters":    Parameters,
		"CareerOptions": CareerOptions,
	})
	if err != nil {
		return "", err
	}

	return doc.ID, err
}
