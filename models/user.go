package models

type User struct {
	Email              string `json:"Email"`
	Age                string `json:"Age"`
	Name               string `json:"Name"`
	SomethingReallyBig string `json:"SomethingReallyBig"`
}
