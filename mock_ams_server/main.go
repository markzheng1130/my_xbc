package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "mx/outbox_service"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedOutboxServiceServer
}

func (s *server) PollAgentEvent(ctx context.Context, in *pb.PollAgentEventRequest) (*pb.PollAgentEventResponse, error) {
	agentEvent1 := pb.AgentEvent{EventId: 1, DeviceId: "11111111-1111-1111-1111-111111111111", EventType: pb.AgentEventType_AGENT_EVENT_REGISTER}
	agentEvent2 := pb.AgentEvent{EventId: 1, DeviceId: "22222222-2222-2222-2222-222222222222", EventType: pb.AgentEventType.AgentEventType_AGENT_EVENT_UPDATE}

	agentEventResponse := pb.PollAgentEventResponse{}
	agentEventResponse.Events = append(agentEventResponse.Events, &agentEvent1)
	agentEventResponse.Events = append(agentEventResponse.Events, &agentEvent2)

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
