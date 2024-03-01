package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "mx/outbox_service"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedOutboxServiceServer
}

func (s *server) PollAgentEvent(ctx context.Context, in *pb.PollAgentEventRequest) (*pb.PollAgentEventResponse, error) {
	log.Printf("pn=[%d], c=[%d]\n", in.GetPartitionNumber(), in.GetCount())

	agentEventResponse := pb.PollAgentEventResponse{}

	for i := 0; i < int(in.GetCount()); i++ {
		agentEvent := pb.AgentEvent{
			EventId:   int64(i),
			DeviceId:  uuid.New().String(),
			EventType: pb.AgentEventType_AGENT_EVENT_UPDATE,
			Contents: string(fmt.Sprintf(`{
				"name": "person-([%d])",
				"contact" : {
				  "address": "taiwan, taipei",
				  "phones": {
					"mobile": "0987-654-321"
				  }
				}
			  }`, i)),
		}

		agentEventResponse.Events = append(agentEventResponse.Events, &agentEvent)
	}

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
