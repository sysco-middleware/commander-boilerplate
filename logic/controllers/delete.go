package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/models"
)

// OnDeleteUser handles a "delete" command
func OnDeleteUser(command *commander.Command) *commander.Event {
	req := models.UserDeleteCommand{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	user := models.UserModel{
		ID: req.ID,
	}

	common.Database.Delete(&user)

	return command.NewEvent("Deleted", 1, command.Key, nil)
}
