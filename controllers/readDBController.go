package controllers

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
	"log"
)

func CountTotalUsers() int {
	// Initialising the global variables for client and context
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Getting the itreator for the db
	iter := client.Collection("Users").Documents(ctx)

	// Looping through the iterator
	var count int = 0
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			return count
		}
		if err != nil {
			log.Fatal(err)
		}
		count++
	}
}

// Make a function to count the total Number of Paid Users
func ListOfUsers() []map[string]interface{} {
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	iter := client.Collection("Users").Documents(ctx)
	var Users []map[string]interface{}

	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Users = append(Users, data.Data())
	}

	return Users
}

// Make a function to Fetch the Data of a particular Person
func ListOfPaidUsers() []map[string]interface{} {
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	iter := client.Collection("Users").Documents(ctx)

	var Users []map[string]interface{}
	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if data.Data()["Paid"] == "false" {
			Users = append(Users, data.Data())
		}
	}

	return Users
}
