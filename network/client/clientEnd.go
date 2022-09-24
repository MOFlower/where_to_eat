package client

import (
	"context"
	"log"
	"strconv"
	pb "where_to_eat/network/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	host string
	port int
}

var (
	ClientEnd = clientEnd{}
	cfg       = config{"127.0.0.1", 2021}
)

type clientEnd struct {
	client pb.W2EClient
}

func init() {
	conn, err := grpc.Dial(cfg.host+strconv.Itoa(cfg.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	ClientEnd.client = pb.NewW2EClient(conn)
}

func (c *clientEnd) BroadCast(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return c.client.BroadCast(ctx, req)
}
