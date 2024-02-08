# Kafka Client

Kafka-Client is a lightweight command-line tool designed to interact with Apache Kafka for testing and development purposes.

## Features

**Create Topics**: Create new topics in your Kafka cluster.
**Publish Messages**: Send messages to a specified Kafka topic.

## Installation

### Prerequisites

- Go (version 1.15 or later)

### Building from Source

Clone the repository and build the binary using the provided Makefile:

```bash
git clone https://github.com/yourusername/kafka-client.git
cd kafka-client
make build
```

## Usage

### Creating Topics

```bash
kafka-client topic create --cluster localhost:9092 --topic new_topic --partitions 3 --replication-factor 1
```

### Publishing Messages

```bash
kafka-client publish --cluster localhost:9092 --topic your_topic --key your_key --value your_message
```
