package topic

import (
	"context"
	"flag"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type CreateCommand struct {
	AdminClientFactory func([]string) *kafka.Client
}

func (c *CreateCommand) Execute(args []string) {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	cluster := createCmd.String("cluster", "localhost:9092", "The Kafka cluster host:port")
	name := createCmd.String("name", "", "The name of the Kafka topic to create")
	partitions := createCmd.Int("partitions", 1, "The number of partitions for the topic")
	replicationFactor := createCmd.Int("replication-factor", 1, "The replication factor for the topic")

	if err := createCmd.Parse(args); err != nil {
		fmt.Println("Error parsing arguments:", err)
		return
	}

	if *name == "" {
		fmt.Println("You must specify a topic name")
		return
	}

	adminClient := c.AdminClientFactory([]string{*cluster})

	ctx := context.Background()
	_, err := adminClient.CreateTopics(ctx, &kafka.CreateTopicsRequest{
		Topics: []kafka.TopicConfig{
			{
				Topic:             *name,
				NumPartitions:     *partitions,
				ReplicationFactor: *replicationFactor,
			},
		},
	})

	if err != nil {
		fmt.Printf("Failed to create topic: %s\n", err)
	} else {
		fmt.Println("Topic created successfully")
	}
}

// PrintCreateUsage prints the usage information for the create topic command.
func PrintCreateUsage() {
	fmt.Println(`Usage of create:
  --cluster string
        The Kafka cluster host:port
  --name string
        The name of the Kafka topic to create
  --partitions int
        The number of partitions for the topic (default 1)
  --replication-factor int
        The replication factor for the topic (default 1)`)
}
