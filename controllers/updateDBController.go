package controllers

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

func UpdateDataToDB(name string, isPaidUser string, stage string) error {
	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding data to the DB
	_, _, err := client.Collection("Users").Add(ctx, map[string]interface{}{
		"First Name": name,
		"Stage":      stage,
		"Paid":       isPaidUser,
	})

	if err != nil {
		return err
	}

	return nil
}
