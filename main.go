package main

import (
	"mx/kafkahelper"
	"time"
)

func main() {
	go kafkahelper.RunConsumer()

	// kafkahelper.Init()
	// go func() {
	// 	for {
	// 		kafkahelper.ProduceEvent("mock_message")
	// 		time.Sleep(1 * time.Second)
	// 	}

	// }()

	for {
		time.Sleep(1 * time.Second)
	}

}
