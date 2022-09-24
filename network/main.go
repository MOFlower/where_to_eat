package main

import (
	"context"
	"log"
	"net"
	"strconv"
	pb "where_to_eat/network/protobuf"
	"where_to_eat/network/server"

	"google.golang.org/grpc"
)

type config struct {
	port int
}

var (
	cfg = config{
		2021,
	}
)

type serverEnd struct {
	pb.UnimplementedW2EServer
	eventCh       chan *pb.Req
	clientEventCh map[int64]chan *pb.Resp
}

func (s *serverEnd) BroadCast(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	s.eventCh <- req
	return server.MergeResp(s.clientEventCh[req.ClientId]), nil
}

func main() {
	s := &serverEnd{}
	s.eventCh = make(chan *pb.Req, 64)

	go server.EventExecDispatcher(s.eventCh, s.clientEventCh)

	listener, err := net.Listen("tcp", strconv.Itoa(cfg.port))
	if err != nil {
		log.Fatal(err.Error())
	}

	gs := grpc.NewServer()
	pb.RegisterW2EServer(gs, s)
	if err := gs.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
