package chatbot

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
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

func ReadQuestion() (map[string][]interface{}, error) {
	// Initialising client and context for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Reading all the documents from the collection
	iter := client.Collection("ChatbotQuestions").Documents(ctx)
	var questions = make(map[string][]interface{})

	for {
		data, err := iter.Next()
		
		if err == iterator.Done {
			break
		}
		
		if err != nil {
			return nil, err
		}
		
		var section string
		section = data.Data()["Section"].(string)
		_, exists := questions[section]
		if !exists {
			questions[section] = []interface{}{
				data.Data(),
			}
			continue
		}else{
			questions[section] = append(questions[section], data.Data())
		}

	}

	return questions, nil
}
