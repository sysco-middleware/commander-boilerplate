package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// UserModel holds the structure of the users table
type UserModel struct {
	ID        *uuid.UUID `gorm:"type:uuid; primary_key"`
	FirstName string     `gorm:"not null;unique"`
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// TableName returns the table name of UserModal
func (model *UserModel) TableName() string {
	return "ServiceUsersView"
}
