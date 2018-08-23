package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/logic/controllers"
	"github.com/sysco-middleware/commander-boilerplate/models"
)

func main() {
	database := common.OpenDatabase()
	commander := common.OpenCommander()

	database.AutoMigrate(&models.UserModel{})

	commander.NewCommandHandle("Create", controllers.OnCreateUser)
	commander.NewCommandHandle("Delete", controllers.OnDeleteUser)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
