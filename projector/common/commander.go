package common

import (
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sysco-middleware/commander"
)

// Commander holds the global commander struct
var Commander *commander.Commander

// OpenCommander opens a new commander interface
func OpenCommander() *commander.Commander {
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	group := os.Getenv("KAFKA_GROUP")

	config := commander.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version = sarama.V1_1_0_0

	Commander = &commander.Commander{
		ConsumerGroup: group,
		Timeout:       5 * time.Second,
		EventTopic:    os.Getenv("COMMANDER_EVENT_TOPIC"),
		CommandTopic:  os.Getenv("COMMANDER_COMMAND_TOPIC"),
	}

	Commander.NewProducer(brokers, config)
	Commander.NewConsumer(brokers, config)

	return Commander
}
