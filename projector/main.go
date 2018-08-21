package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
	"github.com/sysco-middleware/commander-boilerplate/projector/controllers"
	"github.com/sysco-middleware/commander-boilerplate/projector/models"
)

func main() {
	commander := common.OpenCommander()
	database := common.OpenDatabase()

	database.AutoMigrate(&models.Users{})

	commander.NewEventHandle("Created", []int{1}, controllers.OnCreatedUser)
	commander.NewEventHandle("Deleted", []int{1}, controllers.OnDeleteUser)
	commander.NewEventHandle("Updated", []int{1}, controllers.OnUpdateUser)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
