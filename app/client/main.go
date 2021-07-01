package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	pb "kubtest.sbrf.ru/tps/pkg/api"
)

func main() {
	var req pb.JapiRequest
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Get hostname error")
	}
	req.Client = hostname

	s, err := json.Marshal(req)
	z := "\n"
	s = append(s, z...)

	conn, err := net.Dial("tcp", os.Getenv("ADAPTER"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn.Write(s)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		var res pb.JapiResponse
		err := json.Unmarshal([]byte(message), &res)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("C:%16.16s A:%16.16s O:%16.16s Dm:%16.16s Db:%16.16s \n%16.16s", res.Client, res.Adapter, res.Orch, res.Dummy, res.Db, res.Data)
		time.Sleep(1 * time.Second)
	}
}
