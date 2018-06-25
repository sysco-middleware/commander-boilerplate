package main

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/service/controllers"
)

var cmd *commander.Commander

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

	cmd = &commander.Commander{
		Consumer: consumer,
		Producer: producer,
	}

	cmd.CloseOnSIGTERM()
	cmd.StartConsuming()

	cmd.NewCommandHandle("create", controllers.OnCreateUser)
	cmd.NewCommandHandle("update", controllers.OnUpdateUser)
	cmd.NewCommandHandle("delete", controllers.OnDeleteUser)
}
