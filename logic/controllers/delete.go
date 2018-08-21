package controllers

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/logic/models"
)

// DeleteModal used during a "delete" command
type DeleteModal struct {
	ID *uuid.UUID `json:"id"`
}

// OnDeleteUser handles a "delete" command
func OnDeleteUser(command *commander.Command) {
	var data DeleteModal

	err := json.Unmarshal(command.Data, &data)
	if err != nil {
		command.NewError("DataParseError", nil)
		return
	}
	fmt.Println("id:", data.ID)

	user := models.UserModel{
		ID: data.ID,
	}
	common.Database.Delete(&user)

	event := command.NewEvent("Deleted", 1, command.Key, nil)
	common.Commander.ProduceEvent(event)
}
