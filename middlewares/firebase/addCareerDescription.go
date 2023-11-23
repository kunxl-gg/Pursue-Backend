package firebase_middleware

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func AddCareerDescriptionToFirebase(name string, description string, topColleges []types.TopCollge) (string, error) {

	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding data to the DB
	doc, _, err := client.Collection("CareerDescription").Add(ctx, map[string]interface{}{
		"Name": name,
		"Description": description,
		"TopColleges": topColleges,
	})
	if err != nil {
		return "", err
	}

	// If everything goes fine you can add the data to the DB
	return doc.ID, nil

}
