package common

import (
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/sysco-middleware/commander"
)

// Commander holds the global commander struct
var Commander *commander.Commander

// OpenCommander opens a new commander interface
func OpenCommander() *commander.Commander {
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	group := os.Getenv("KAFKA_GROUP")

	producerConfig := sarama.NewConfig()
	consumerConfig := cluster.NewConfig()

	producerConfig.Version = sarama.V1_1_0_0
	consumerConfig.Version = sarama.V1_1_0_0

	Commander = &commander.Commander{
		ConsumerGroup: group,
		Timeout:       5 * time.Second,
		EventTopic:    os.Getenv("COMMANDER_EVENT_TOPIC"),
		CommandTopic:  os.Getenv("COMMANDER_COMMAND_TOPIC"),
	}

	Commander.NewProducer(brokers, producerConfig)
	Commander.NewConsumer(brokers, consumerConfig)

	return Commander
}
