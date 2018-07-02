package controllers

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander"
)

// CreateModal used during a "create" command
type CreateModal struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// OnCreateUser handles a "create" command
func OnCreateUser(command *commander.Command) {
	var data CreateModal

	UnmarshalErr := json.Unmarshal(command.Data, &data)
	if UnmarshalErr != nil {
		command.NewError("DataParseError", nil)
		return
	}

	// ...

	key := uuid.NewV4()
	res, MarshalErr := json.Marshal(data)

	if MarshalErr != nil {
		command.NewError("ResponseParseError", nil)
		return
	}

	command.NewEvent("Created", key, res)
}
