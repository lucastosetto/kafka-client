package publisher

import (
	"context"
	"flag"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter interface {
	WriteMessages(context.Context, ...kafka.Message) error
	Close() error
}

type PublishCommand struct {
	WriterFactory func(config kafka.WriterConfig) KafkaWriter
}

func (p *PublishCommand) Execute(args []string) {
	publishCmd := flag.NewFlagSet("publish", flag.ExitOnError)
	cluster := publishCmd.String("cluster", "localhost:9092", "The Kafka cluster host:port")
	topic := publishCmd.String("topic", "", "The Kafka topic to publish to")
	key := publishCmd.String("key", "", "The key for the message")
	value := publishCmd.String("value", "", "The value of the message")

	if err := publishCmd.Parse(args); err != nil {
		fmt.Println("Error parsing arguments:", err)
		return
	}

	clusterVal := *cluster
	topicVal := *topic
	keyVal := *key
	valueVal := *value

	writer := p.WriterFactory(kafka.WriterConfig{
		Brokers: []string{clusterVal},
		Topic:   topicVal,
	})
	defer writer.Close()

	msg := kafka.Message{
		Key:   []byte(keyVal),
		Value: []byte(valueVal),
	}

	ctx := context.Background()

	err := writer.WriteMessages(ctx, msg)
	if err != nil {
		fmt.Printf("Failed to publish message: %s\n", err)
	} else {
		fmt.Println("Message published successfully")
	}
}

// PrintPublishUsage prints the usage information for the publish command.
func PrintPublishUsage() {
	fmt.Println(`Usage of publish:
  --cluster string
        The Kafka cluster host:port (default "localhost:9092")
  --topic string
        The Kafka topic to publish to
  --key string
        The key for the message
  --value string
        The value of the message`)
}
