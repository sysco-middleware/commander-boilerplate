package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
	"github.com/sysco-middleware/commander-boilerplate/projector/controllers"
	"github.com/sysco-middleware/commander-boilerplate/projector/models"
)

func main() {
	commander := common.OpenCommander()
	database := common.OpenDatabase()

	// Let's set the offset of the projector to 0
	go func() {
		for event := range commander.Events() {
			switch message := event.(type) {
			case kafka.AssignedPartitions:
				parts := make([]kafka.TopicPartition, 0, len(message.Partitions))
				for _, part := range message.Partitions {
					part.Offset = kafka.Offset(0)
					parts = append(parts, part)
				}

				commander.Consumer.Assign(parts)
			case kafka.RevokedPartitions:
				commander.Consumer.Unassign()
			}
		}
	}()

	database.AutoMigrate(&models.Users{})

	commander.NewEventHandle("Created", controllers.OnCreatedUser)
	commander.NewEventHandle("Deleted", controllers.OnDeleteUser)
	commander.NewEventHandle("Updated", controllers.OnUpdateUser)

	commander.CloseOnSIGTERM()
	commander.StartConsuming()
}
