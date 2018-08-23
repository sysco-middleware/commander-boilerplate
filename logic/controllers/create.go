package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/models"
)

// OnCreateUser handles a "create" command
func OnCreateUser(command *commander.Command) *commander.Event {
	req := models.UserCreatedCommand{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	user := models.UserModel{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	query := common.Database.Save(&user)
	if query.Error != nil {
		return command.NewErrorEvent("UserExists", nil)
	}

	event := models.UserCreatedEvent{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	res, _ := json.Marshal(event)
	return command.NewEvent("Created", 1, *user.ID, res)
}
