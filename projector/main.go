package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/projector/common"
	"github.com/sysco-middleware/commander-boilerplate/projector/controllers"
	"github.com/sysco-middleware/commander-boilerplate/projector/models"
)

func main() {
	common.OpenCommander()
	common.OpenDatabase()

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

	common.Database.AutoMigrate(&models.Users{})

	common.Commander.NewEventHandle("Created", controllers.OnCreatedUser)
	common.Commander.NewEventHandle("Deleted", controllers.OnDeleteUser)
	common.Commander.NewEventHandle("Updated", controllers.OnUpdateUser)

	common.Commander.CloseOnSIGTERM()
	common.Commander.StartConsuming()
}
