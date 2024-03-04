package kafkahelper

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	run     = true
	groupId = "my_group_5"
)

func RunConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Fatalf("Init kafka consumer client encountered exceptions: %v", err)
	}

	err = c.SubscribeTopics(topicList, nil)

	for run == true {
		ev := c.Poll(100)
		switch e := ev.(type) {

		case *kafka.Message:
			log.Printf("[Received message][%v]", string(e.Value))

		case kafka.Error:
			log.Printf("%% Error: %v\n", e)
			run = false
		}
	}

	c.Close()
}
