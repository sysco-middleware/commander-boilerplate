package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
)

// DeleteModal used during a "delete" command
type DeleteModal struct {
	User string `json:"user"`
}

// OnDeleteUser handles a "delete" command
func OnDeleteUser(command *commander.Command) {
	var data DeleteModal

	UnmarshalErr := json.Unmarshal(command.Data, &data)
	if UnmarshalErr != nil {
		command.NewError("DataParseError", nil)
		return
	}

	// ...

	event := command.NewEvent("Delete", command.Key, nil)
	common.Commander.ProduceEvent(event)
}
