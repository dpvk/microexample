package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	pb "kubtest.sbrf.ru/tps/pkg/api"
)

func main() {
	// Start listener
	ln, _ := net.Listen("tcp", os.Getenv("LISTEN"))

	// Connect to
	resolver.SetDefaultScheme("dns")
	conn, err := grpc.Dial(os.Getenv("ORCH"), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("Launching server...")
	for {
		{
			clnt, err := ln.Accept()
			if err != nil {
				log.Fatalf("Listen error: %v", err)
			}
			defer clnt.Close()
			for {

				message, _ := bufio.NewReader(clnt).ReadString('\n')
				fmt.Print("Message Received:", string(message))
				var req pb.JapiRequest
				err := json.Unmarshal([]byte(message), &req)
				if err != nil {
					log.Printf("json parse error: %v", err)
					//clnt.Close()
					break
				}

				c := pb.NewAdapterClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				r, err := c.Send(ctx, &pb.Request{Client: req.Client})
				if err != nil {
					log.Printf("gRPC request error: %v", err)
					continue
				}
				log.Printf("Received:\n Clnt: %v \n Orch:  %v \n Dumm: %v \n DB  : %v \n Data:%v", r.GetClient(), r.GetOrch(), r.GetDummy(), r.GetDb(), r.GetData())

				hostname, _ := os.Hostname()
				res := pb.JapiResponse{Client: r.GetClient(), Adapter: hostname, Orch: r.GetOrch(), Dummy: r.GetDummy(), Db: r.GetDb(), Data: r.GetData()}

				s, _ := json.Marshal(res)

				s = append(s, "\n"...)

				clnt.Write(s)
			}
		}
	}

}
