package main

import (
	"mx/kafkahelper"
	"time"
)

func main() {
	go kafkahelper.RunConsumer()

	// go func() {
	// 	for true {
	// 		kafkahelper.RunProducer()
	// 		time.Sleep(10 * time.Second)
	// 	}

	// }()

	for true {
		time.Sleep(1 * time.Second)
	}

}
