package kafkahelper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	run     = true
	groupId = "my_group_1"
)

func RunConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Fatalf("Init kafka consumer client encountered exceptions: %v", err)
	}

	_ = c.SubscribeTopics(topicList, nil)

	for run {
		ev := c.Poll(100)
		switch e := ev.(type) {

		case *kafka.Message:
			var m map[string]interface{}
			if err := json.Unmarshal(e.Value, &m); err != nil {
				fmt.Printf("Error: %v\n\n", err)
			}

			log.Printf("[Received message]%v", string(e.Value))
			// log.Printf("[Received message]%v, %v", e.Headers, m)

		case kafka.Error:
			log.Printf("%% Error: %v\n", e)
			run = false
		}
	}

	c.Close()
}
