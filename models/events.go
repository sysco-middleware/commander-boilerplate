package models

import uuid "github.com/satori/go.uuid"

// UserCreatedEvent holds the content of a "created" event
type UserCreatedEvent struct {
	ID        *uuid.UUID `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
}
