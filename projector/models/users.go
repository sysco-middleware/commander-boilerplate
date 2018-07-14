package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Users gorm database table struct
type Users struct {
	ID        *uuid.UUID `gorm:"type:uuid; primary_key"`
	FirstName string     `gorm:"not null;unique"`
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// TableName table name of User
func (u *Users) TableName() string {
	return "ProjectorUsersView"
}
