package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/logic/controllers"
)

func main() {
	common.OpenDatabase()
	common.OpenCommander()

	common.Database.AutoMigrate(&controllers.UserModel{})

	common.Commander.NewCommandHandle("Create", controllers.OnCreateUser)
	common.Commander.NewCommandHandle("Update", controllers.OnUpdateUser)
	common.Commander.NewCommandHandle("Delete", controllers.OnDeleteUser)

	common.Commander.CloseOnSIGTERM()
	common.Commander.StartConsuming()
}
