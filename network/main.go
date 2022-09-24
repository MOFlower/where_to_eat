package main

import (
	"log"
	"net"
	pb "where_to_eat/network/protobuf"
	"where_to_eat/network/server"

	"google.golang.org/grpc"
)

var addr = "127.0.0.1:2021"

func main() {
	println("hello world! start server...")

	s := server.NewServerEnd()

	go server.EventExecDispatcher(s.ReqCh, s.ClientRespChMap)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	gs := grpc.NewServer()
	pb.RegisterW2EServer(gs, &s)
	if err := gs.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
