package main

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/service/common"
	"github.com/sysco-middleware/commander-boilerplate/service/controllers"
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

	common.Commander = &commander.Commander{
		Consumer: consumer,
		Producer: producer,
	}

	common.Commander.NewCommandHandle("Create", controllers.OnCreateUser)
	common.Commander.NewCommandHandle("Update", controllers.OnUpdateUser)
	common.Commander.NewCommandHandle("Delete", controllers.OnDeleteUser)

	common.Commander.CloseOnSIGTERM()
	common.Commander.StartConsuming()
}
