package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	projectID        = "cascade-masters"
	topicID          = "dev-example-topic"
	subscriptionName = "your-subscription-name"
)


func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	topic := client.Topic(topicID)
	defer topic.Stop()

	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Hello from Producer! Message #%d", i)
		result := topic.Publish(ctx, &pubsub.Message{
			Data: []byte(msg),
		})
		id, err := result.Get(ctx)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
			continue
		}
		fmt.Printf("Published message %d with ID: %s\n", i, id)
		time.Sleep(1 * time.Second)
	}
}
