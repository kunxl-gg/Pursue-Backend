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

func AddUserChoicesInRepository(ID string, DatabaseTitle string, Parameters []string, CareerOptions []string) (string, error) {
	// Initialising ctx and client for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding Entry to Firebase
	_, err := client.Collection(DatabaseTitle).Doc(ID).Set(ctx, map[string]interface{}{
		"Parameters":    Parameters,
		"CareerOptions": CareerOptions,
	})
	if err != nil {
		return "", err
	}

	return "Added an Entry to Repository", err
}

func DeleteRepository(ID string, DatabaseTable string) (string, error) {
	// Initialising firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Deleting the entry from firebase
	_, err := client.Collection(DatabaseTable).Doc(ID).Delete(ctx)
	if err != nil {
		return "", nil
	}

	return "Deleted Item From Repository Successfully", nil
}
