package models

// UserCreatedCommand holds the content of a "created" command
type UserCreatedCommand struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
