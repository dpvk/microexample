package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	pb "kubtest.sbrf.ru/tps/pkg/api"
)

type services struct {
	Dummy *grpc.ClientConn
	Db    *grpc.ClientConn
}

var svcctx services

type server struct {
	pb.UnimplementedAdapterServer
}

func (s *server) Send(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("From client %v", in.Client)
	// Local routine
	hostname, _ := os.Hostname()
	hostname = hostname + ":" + strconv.Itoa(os.Getpid())
	resp := pb.Response{Client: in.Client, Orch: hostname}

	// SVC_Dummy routine
	dummy, err := RequestDummy(in.Client, hostname)
	if err != nil {
		return &resp, nil
	}
	resp.Dummy = dummy
	// SVC_DB routine
	db, data, err := RequestDB(in.Client, hostname)
	if err != nil {
		return &resp, nil
	}
	resp.Db = db
	resp.Data = data
	return &resp, nil
}

func main() {
	resolver.SetDefaultScheme("dns")
	svcctx.Dummy, _ = grpc.Dial(os.Getenv("DUMMY"), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	defer svcctx.Dummy.Close()

	svcctx.Db, _ = grpc.Dial(os.Getenv("DB"), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	defer svcctx.Db.Close()

	port := os.Getenv("LISTEN")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdapterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RequestDummy(client string, orch string) (string, error) {

	c := pb.NewDummyClient(svcctx.Dummy)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Process(ctx, &pb.DummyRequest{Client: client, Orch: orch})
	if err != nil {
		log.Printf("Dummy process error: %v", err)
		return "", err

	}

	return r.GetDummy(), nil
}

func RequestDB(client string, orch string) (string, string, error) {

	c := pb.NewDBSClient(svcctx.Db)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Process(ctx, &pb.DBRequest{Client: client, Orch: orch})
	if err != nil {
		log.Printf("Db process error: %v", err)
		return "", "", err

	}

	return r.GetDb(), r.GetData(), nil
}
