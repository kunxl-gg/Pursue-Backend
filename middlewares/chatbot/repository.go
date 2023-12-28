package chatbot

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

func AddQuestionInRepository(Questions []string, Answers []string) (string, error) {
	// Initialising client and context
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding a new object to firebase Collection
	doc, _, err := client.Collection("Repository").Add(ctx, map[string]interface{}{
		"Options":      Questions,
		"CareerChoice": Answers,
	})
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}

func AddUserChoicesInRepository(CareerOptions []string, Parameters []string) (string, error) {
	// Initialising ctx and client for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding the Entry to Firebase
	doc, _, err := client.Collection("Repository").Add(ctx, map[string]interface{}{
		"Parameters":    Parameters,
		"CareerOptions": CareerOptions,
	})
	if err != nil {
		return "", err
	}

	return doc.ID, err
}

func ReadUserChoicesWithOneFixedVariable(CollectionName string) ([]map[string]interface{}, error) {
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Reading from the firebase collection
	iter := client.Collection(CollectionName).Documents(ctx)
	var choice []map[string]interface{}

	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		choice = append(choice, data.Data())
	}

	return choice, nil
}
