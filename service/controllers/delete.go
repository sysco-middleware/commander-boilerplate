package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
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

	command.NewEvent("Delete", command.Key, nil)
}
