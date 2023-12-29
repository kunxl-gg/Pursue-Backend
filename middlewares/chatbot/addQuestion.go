package chatbot

import (
	"fmt"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

func AddStudentQuestion(Section string, Question string, Option []string) (string, error) {
	// Initialising client and context for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding Question to Firebase
	doc, _, err := client.Collection("ChatbotQuestions").Add(ctx, map[string]interface{}{
		"Section":  Section,
		"Question": Question,
		"Options":  Option,
	})
	if err != nil {
		return "", err
	}

	// If everything goes fine send a success Message
	return doc.ID, nil
}

func ReadQuestion() ([]map[string]interface{}, error) {
	// Initialising client and context for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Reading all the documents from the collection
	iter := client.Collection("ChatbotQuestions").Documents(ctx)
	var _ []map[string]interface{}
	fmt.Print(iter)
	return nil, nil
}
