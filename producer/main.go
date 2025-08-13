package main

import (
	"context"
	"fmt"
	"log"
	"time"
"github.com/Niko98108/go-event-driven-architecture/constant"
	"cloud.google.com/go/pubsub"
)


func main() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, )
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
