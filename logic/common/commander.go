package common

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sysco-middleware/commander"
)

// Commander holds the global commander struct
var Commander *commander.Commander

// OpenCommander opens a new commander interface
func OpenCommander() *commander.Commander {
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

	Commander = &commander.Commander{
		Consumer: consumer,
		Producer: producer,
	}

	return Commander
}
