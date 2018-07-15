package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
	"github.com/sysco-middleware/commander-boilerplate/projector/models"
)

// OnCreatedUser handles a "created" event
func OnCreatedUser(command *commander.Event) {
	var user models.Users

	dataParseError := json.Unmarshal(command.Data, &user)
	// Parse the data back to a struct
	if dataParseError != nil {
		// Optionally could you panic on a error
		// panic(dataParseError)
		return
	}

	query := common.Database.Create(&user)
	// A user already exists if a error occures
	if query.Error != nil {
		// Optionally could you panic on a error
		// panic(query.Error)
		return
	}
}
