package server

import (
	"context"
	pb "where_to_eat/network/protobuf"
)

type ServerEnd struct {
	pb.UnimplementedW2EServer
	ReqCh           chan *pb.Req
	ClientRespChMap map[int64]chan *pb.Resp
}

func NewServerEnd() ServerEnd {
	ret := ServerEnd{}
	ret.ReqCh = make(chan *pb.Req, 64)
	return ret
}

func (s *ServerEnd) Pull(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	ret := &pb.Resp{}
	select {
	case ret = <-s.ClientRespChMap[req.ClientId]:
	default:
	}
	return ret, nil
}

func (s *ServerEnd) Push(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	s.ReqCh <- req
	return &pb.Resp{}, nil
}
