package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/models"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
)

// OnCreatedUser handles a "created" event
func OnCreatedUser(event *commander.Event) {
	req := models.UserCreatedEvent{}
	err := json.Unmarshal(event.Data, &req)
	if err != nil {
		return
	}

	user := models.UserView{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	query := common.Database.Save(&user)
	if query.Error != nil {
		return
	}
}
