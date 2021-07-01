package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	pb "kubtest.sbrf.ru/tps/pkg/api"
)

type server struct {
	pb.UnimplementedDummyServer
}

func (s *server) Process(ctx context.Context, in *pb.DummyRequest) (*pb.DumyResponse, error) {
	log.Printf("Got request: client %v orch %v", in.Client, in.Orch)
	// Local routine
	hostname, _ := os.Hostname()
	hostname = hostname + ":" + strconv.Itoa(os.Getpid())

	resp := pb.DumyResponse{Dummy: hostname}
	return &resp, nil
}

func main() {
	port := os.Getenv("LISTEN")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDummyServer(s, &server{})
	log.Printf("Dummy listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
