package controllers

import (
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/models"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
)

// OnDeleteUser handles a "delete" event
func OnDeleteUser(event *commander.Event) {
	user := models.UserView{
		ID: &event.Key,
	}

	query := common.Database.Delete(&user)
	if query.Error != nil {
		return
	}
}
