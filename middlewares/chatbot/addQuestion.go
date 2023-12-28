package chatbot

import (
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
