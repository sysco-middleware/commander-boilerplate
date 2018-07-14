package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
	"github.com/sysco-middleware/commander-boilerplate/projector/controllers"
	"github.com/sysco-middleware/commander-boilerplate/projector/models"
)

func main() {
	servers := os.Getenv("KAFKA_SERVERS")
	group := os.Getenv("KAFKA_GROUP")

	producer := commander.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
	})

	consumer := commander.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               servers,
		"group.id":                        group,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	common.Commander = &commander.Commander{
		Consumer: consumer,
		Producer: producer,
	}

	// Let's set the offset of the projector to 0
	go func() {
		for event := range common.Commander.Events() {
			switch message := event.(type) {
			case kafka.AssignedPartitions:
				parts := make([]kafka.TopicPartition, 0, len(message.Partitions))
				for _, part := range message.Partitions {
					part.Offset = kafka.Offset(0)
					parts = append(parts, part)
				}

				common.Commander.Consumer.Assign(parts)
			case kafka.RevokedPartitions:
				common.Commander.Consumer.Unassign()
			}
		}
	}()

	common.Database = OpenDatabase()
	common.Database.AutoMigrate(&models.Users{})

	common.Commander.NewEventHandle("Created", controllers.OnCreatedUser)
	common.Commander.NewEventHandle("Deleted", controllers.OnDeleteUser)
	common.Commander.NewEventHandle("Updated", controllers.OnUpdateUser)

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
