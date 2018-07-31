package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/logic/controllers"
)

func main() {
	database := common.OpenDatabase()
	commander := common.OpenCommander()

	database.AutoMigrate(&controllers.UserModel{})

	commander.NewCommandHandle("Create", controllers.OnCreateUser)
	commander.NewCommandHandle("Update", controllers.OnUpdateUser)
	commander.NewCommandHandle("Delete", controllers.OnDeleteUser)

	commander.CloseOnSIGTERM()
	commander.StartConsuming()
}
