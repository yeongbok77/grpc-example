package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}
	grpcServer := grpc.NewServer() // 生成一个新的grpc服务

}
