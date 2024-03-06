package kafkahelper

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	bootstrapServers = "localhost:9092"
	topicList        = []string{"xbc-agent-event", "xbc-agent-internal-activity"}
	producerClient   *kafka.Producer
)

func Init() {
	var err error
	producerClient, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Init kafka producer client encountered exceptions: %v", err)
	}
}

func ProduceEvent(message string) {

	for _, topic := range topicList {
		err := producerClient.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)

		if err != nil {
			log.Fatalf("When kafka producer client producing message encountered exceptions: %v", err)
		}
	}

	go func() {
		for e := range producerClient.Events() {
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

	producerClient.Flush(15 * 1000)
}
