package models

import (
	"time"

	"github.com/jinzhu/gorm"
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
func (user *UserModel) TableName() string {
	return "ServiceUsersView"
}

// BeforeCreate generates required data before creating
func (user *UserModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// UserView gorm database table struct
type UserView struct {
	ID        *uuid.UUID `gorm:"type:uuid; primary_key"`
	FirstName string     `gorm:"not null;unique"`
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// TableName table name of User
func (user *UserView) TableName() string {
	return "ProjectorUsersView"
}
