package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "mx/xbc_v2"

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
			DeviceId:  "11111111-1111-1111-1111-111111111111",
			EventType: pb.AgentEventType_AGENT_EVENT_UPDATE,
			Contents:  "{\"os_arch\": \"x86_64\", \"platform\": \"win32\", \"products\": [{\"company_id\": \"706202DE-9CF0-4704-A1B7-70DCAD85E690\", \"product_id\": \"96529838-586c-4e0b-87dc-607443b76981\"}], \"trace_id\": \"b109c5ed-faf1-4ca6-a4fe-d8e13aef65d8\", \"device_id\": \"56e173e9-8542-f15b-e531-d7ded8fc920b\", \"wsc_version\": \"1.1.0.3695\", \"ceta_version\": \"1.1.0.888\", \"agent_version\": \"1.1.0.3970\", \"is_base_image\": false, \"xdr_device_id\": \"56e173e9-8542-f15b-e531-d7ded8fc920b\"}",
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
