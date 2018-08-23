package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/models"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
	"github.com/sysco-middleware/commander-boilerplate/projector/controllers"
)

func main() {
	commander := common.OpenCommander()
	database := common.OpenDatabase()

	database.AutoMigrate(&models.UserView{})

	commander.NewEventHandle("Created", []int{1}, controllers.OnCreatedUser)
	commander.NewEventHandle("Deleted", []int{1}, controllers.OnDeleteUser)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
