package models

import uuid "github.com/satori/go.uuid"

// UserCreatedCommand holds the content of a "created" command
type UserCreatedCommand struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserDeleteCommand holds the content of a "delete" command
type UserDeleteCommand struct {
	ID *uuid.UUID `json:"id"`
}
