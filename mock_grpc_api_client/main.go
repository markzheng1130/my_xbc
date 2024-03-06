package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	pb "mx/xbc_v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr            = flag.String("addr", "localhost:50051", "the address to connect to")
	partitionNumber = flag.Int64("pn", 0, "A table partition's number")
	count           = flag.Int64("c", 0, "Max agent event count")
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

	// step1, invoke poll agent event API.
	r, err := c.PollAgentEvent(ctx, &pb.PollAgentEventRequest{PartitionNumber: *partitionNumber, Count: *count})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for _, agentEvent := range r.GetEvents() {
		var contents map[string]interface{}
		if err := json.Unmarshal([]byte(agentEvent.Contents), &contents); err != nil {
			fmt.Printf("Error: %v\n\n", err)
		}

		fmt.Printf("[%v][%v][%v][%v]\n", agentEvent.EventId, agentEvent.DeviceId, agentEvent.EventType, contents)
	}

	// step2, invoke commit API.
	offsetId := count
	consumerName := "mockOFS"

	r2, err := c.CommitAgentEvent(ctx, &pb.CommitAgentEventRequest{PartitionNumber: *partitionNumber, OffsetId: *offsetId, ConsumerName: consumerName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("[Commit result][%v][%T]\n", r2, r2)
}

// go run ./mock_grpc_api_client/main.go --pn=2 -c=5
// sudo lsof -i:50051
