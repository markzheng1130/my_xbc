package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "mx/helloworld"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) PollAgentEvent(ctx context.Context, in *pb.PollAgentEventRequest) (*pb.PollAgentEventResponse, error) {
	agentEvent1 := pb.AgentEvent{event_id: 1, device_id: "11111111-1111-1111-1111-111111111111", event_type: pb.AgentEventType.AGENT_EVENT_REGISTER}
	agentEvent2 := pb.AgentEvent{event_id: 1, device_id: "22222222-2222-2222-2222-222222222222", event_type: pb.AgentEventType.AGENT_EVENT_UPDATE}

	agentEventResponse := pb.PollAgentEventResponse{}
	agentEventResponse.Events = append(agentEventResponse.Events, agentEvent1, agentEvent2)

	return &agentEventResponse, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOutboxServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
