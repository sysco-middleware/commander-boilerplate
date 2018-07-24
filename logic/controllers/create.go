package controllers

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
)

// CreateModal is used during a "create" command
type CreateModal struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

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
func (modal *UserModel) TableName() string {
	return "ServiceUsersView"
}

// OnCreateUser handles a "create" command
func OnCreateUser(command *commander.Command) {
	var data CreateModal

	// Parse the data back to a struct
	UnmarshalErr := json.Unmarshal(command.Data, &data)
	if UnmarshalErr != nil {
		command.NewError("DataParseError", nil)
		return
	}

	// Prepare a new user
	key := uuid.NewV4()
	user := UserModel{
		ID:        &key,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	query := common.Database.Create(&user)
	// A user already exists if a error occures
	if query.Error != nil {
		event := command.NewError("UserExists", nil)
		common.Commander.ProduceEvent(event)
		return
	}

	res, MarshalErr := json.Marshal(user)

	if MarshalErr != nil {
		command.NewError("ResponseParseError", nil)
		return
	}

	event := command.NewEvent("Created", key, res)
	common.Commander.ProduceEvent(event)
}
