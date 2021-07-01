package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	pb "kubtest.sbrf.ru/tps/pkg/api"
)

var gdbconn *sql.DB

type server struct {
	pb.UnimplementedDBSServer
}

func (s *server) Process(ctx context.Context, in *pb.DBRequest) (*pb.DBResponse, error) {
	log.Printf("DB service. request: client %v orch %v", in.Client, in.Orch)
	// Local routine
	hostname, _ := os.Hostname()
	hostname = hostname + ":" + strconv.Itoa(os.Getpid())

	var userid int
	err := gdbconn.QueryRow(`insert into table1 (val) values ('test text1') returning Id`).Scan(&userid)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	resp := pb.DBResponse{Db: hostname, Data: "ID From DB " + strconv.Itoa(userid)}
	return &resp, nil
}

func main() {

	dbconn, err := sql.Open("postgres", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()
	gdbconn = dbconn

	log.Print("Connected to database")

	lis, err := net.Listen("tcp", os.Getenv("LISTEN"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDBSServer(s, &server{})
	log.Printf("DBS listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
