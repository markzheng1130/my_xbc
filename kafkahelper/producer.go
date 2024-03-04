package kafkahelper

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	bootstrapServers = "localhost:9092"
	topicList        = []string{"my_topic_1", "my_topic_2"}
	partitionList    = []int{1, 2}
)

func RunProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Init kafka producer client encountered exceptions: %v", err)
	}

	defer p.Close()

	for i, topic := range topicList {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(partitionList[i])},
			Value:          []byte("mock_message"),
		}, nil)

		if err != nil {
			log.Fatalf("When kafka producer client producing message encountered exceptions: %v", err)
		}
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	p.Flush(15 * 1000)
}
