package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "mx/outbox_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr            = flag.String("addr", "localhost:50051", "the address to connect to")
	partitionNumber = int64(2)
	count           = int64(5000)
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewOutboxServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.PollAgentEvent(ctx, &pb.PollAgentEventRequest{PartitionNumber: partitionNumber, Count: count})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, agentEvent := range r.GetEvents() {
		fmt.Printf("[Get agent event][%v]\n\n", agentEvent)
	}
}
