package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "mx/xbc_v2"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedOutboxServiceServer
}

func (s *server) PollAgentEvent(ctx context.Context, in *pb.PollAgentEventRequest) (*pb.PollAgentEventResponse, error) {
	log.Printf("[Received][Poll API][partition number][%d][count][%d]\n", in.GetPartitionNumber(), in.GetCount())

	agentEventResponse := pb.PollAgentEventResponse{}

	for i := 0; i < int(in.GetCount()); i++ {
		agentEvent := pb.AgentEvent{
			EventId:   time.Now().UnixNano(),
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

func (s *server) CommitAgentEvent(ctx context.Context, in *pb.CommitAgentEventRequest) (*emptypb.Empty, error) {
	log.Printf("[Received][Commit API][partition number][%v][offset id][%v][consumer name][%v]\n", in.PartitionNumber, in.OffsetId, in.ConsumerName)
	return &emptypb.Empty{}, nil
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
