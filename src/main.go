package main

import (
	"fmt"
	"os"

	"github.com/lucastosetto/kafka-client/src/publisher"
	"github.com/lucastosetto/kafka-client/src/topic"
	"github.com/segmentio/kafka-go"
)

func printGeneralHelp() {
	fmt.Println(`Usage: kafka-client <command> [subcommand] [options]

Commands:
	topic       Manage topics (use "kafka-client topic --help" for more)
	publish     Publish messages to a Kafka topic

Use "kafka-client <command> --help" for more information about a command.`)
}

func printTopicHelp() {
	fmt.Println(`Usage: kafka-client topic <subcommand> [options]

Subcommands:
	create     Create a new Kafka topic

Use "kafka-client topic <subcommand> --help" for subcommand options.`)
}

func main() {
	if len(os.Args) < 2 {
		printGeneralHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "topic":
		if len(os.Args) < 3 || os.Args[2] == "--help" {
			printTopicHelp()
			return
		}
		handleTopicCommand(os.Args[2:])
	case "publish":
		handlePublishCommand(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printGeneralHelp()
	}
}

func handlePublishCommand(args []string) {
	if len(args) > 0 && (args[0] == "-h" || args[0] == "--help") {
		publisher.PrintPublishUsage()
		return
	}
	command := &publisher.PublishCommand{
		WriterFactory: func(config kafka.WriterConfig) publisher.KafkaWriter {
			return kafka.NewWriter(config)
		},
	}
	command.Execute(args)
}

func handleTopicCommand(args []string) {
	if len(args) < 1 {
		printTopicHelp()
		return
	}

	subcommand := args[0]
	switch subcommand {
	case "create":
		if len(args) > 1 && (args[1] == "-h" || args[1] == "--help") {
			topic.PrintCreateUsage()
			return
		}
		command := &topic.CreateCommand{
			AdminClientFactory: func(brokers []string) *kafka.Client {
				return &kafka.Client{Addr: kafka.TCP(brokers[0])}
			},
		}
		command.Execute(args[1:])
	default:
		fmt.Printf("Unknown topic subcommand: %s\n", subcommand)
		printTopicHelp()
	}
}
