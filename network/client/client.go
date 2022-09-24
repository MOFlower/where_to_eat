package client

import (
	"context"
	"log"
	pb "where_to_eat/network/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = "127.0.0.1:2021"

	ClientEnd = clientEnd{}
)

type clientEnd struct {
	client pb.W2EClient
	p      *persister
}

func init() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	ClientEnd.client = pb.NewW2EClient(conn)
	ClientEnd.p = &iPersister
}

func (c *clientEnd) Pull(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return c.client.Pull(ctx, req)
}

func (c *clientEnd) Push(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return c.client.Push(ctx, req)
}

func (c *clientEnd) Commit(ctx context.Context, msg string) error {
	c.p.Append([]byte(msg))
	return nil
}
