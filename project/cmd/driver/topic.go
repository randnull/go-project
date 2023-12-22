package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	ctx := context.Background()

	name := "getter"
	partitions := 1
	replicas := 1

	conn, err := kafka.DialContext(ctx, "tcp", "localhost:29092")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             name,
		NumPartitions:     partitions,
		ReplicationFactor: replicas,
	})
	if err != nil {
		log.Fatal(err)
	}

	name = "setter"
	partitions = 1
	replicas = 1

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             name,
		NumPartitions:     partitions,
		ReplicationFactor: replicas,
	})
	if err != nil {
		log.Fatal(err)
	}

}
