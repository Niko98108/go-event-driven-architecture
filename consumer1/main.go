package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

const (
	projectID        = "cascade-masters"
	subscriptionName = "dev-example-consumer-one-sub"
)

func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subscriptionName)
	sub.ReceiveSettings.Synchronous = false
	sub.ReceiveSettings.MaxOutstandingMessages = 10

	fmt.Println("Consumer is listening for messages...")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(m.Data))
		m.Ack()
	})
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}
