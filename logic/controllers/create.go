package controllers

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/logic/models"
)

// CreateModel is used during a "create" command
type CreateModel struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// OnCreateUser handles a "create" command
func OnCreateUser(command *commander.Command) {
	var data CreateModel

	// Parse the data back to a struct
	err := json.Unmarshal(command.Data, &data)
	if err != nil {
		command.NewError("DataParseError", nil)
		return
	}

	// Prepare a new user
	key := uuid.NewV4()
	user := models.UserModel{
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

	res, _ := json.Marshal(user)
	event := command.NewEvent("Created", 1, key, res)
	common.Commander.ProduceEvent(event)
}
