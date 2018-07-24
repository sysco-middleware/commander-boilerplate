package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/logic/common"
	"github.com/sysco-middleware/commander-boilerplate/logic/controllers"
)

func main() {
	servers := os.Getenv("KAFKA_SERVERS")
	group := os.Getenv("KAFKA_GROUP")

	producer := commander.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
	})

	consumer := commander.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    servers,
		"group.id":             group,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	common.Database = OpenDatabase()
	common.Commander = &commander.Commander{
		Consumer: consumer,
		Producer: producer,
	}

	common.Database.AutoMigrate(&controllers.UserModel{})

	common.Commander.NewCommandHandle("Create", controllers.OnCreateUser)
	common.Commander.NewCommandHandle("Update", controllers.OnUpdateUser)
	common.Commander.NewCommandHandle("Delete", controllers.OnDeleteUser)

	common.Commander.CloseOnSIGTERM()
	common.Commander.StartConsuming()
}

// OpenDatabase opens a new database connection
func OpenDatabase() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")

	options := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)
	db, err := gorm.Open("postgres", options)

	if err != nil {
		panic(err)
	}

	return db
}
