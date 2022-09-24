package server

import (
	"log"
	pb "where_to_eat/network/protobuf"
)

func EventExecDispatcher(in chan *pb.Req, m map[int64]chan *pb.Resp) {
	event := <-in
	if _, ok := m[event.ClientId]; !ok {
		m[event.ClientId] = make(chan *pb.Resp, 16)
	}
	switch event.Cmd {
	case pb.CommandType_APPEND_LOG:
		go appendHandler(event, m)
	default:
		log.Fatalln("no found cmd")
	}
}

func appendHandler(req *pb.Req, m map[int64]chan *pb.Resp) {
	for k := range m {
		if k == req.ClientId {
			continue
		}

		resp := &pb.Resp{}
		resp.ToClientId = k
		resp.Msg = req.Msg
		m[k] <- resp
	}
}

func MergeResp(respCh chan *pb.Resp) *pb.Resp {
	ret := &pb.Resp{}
	for {
		select {
		case resp := <-respCh:
			switch resp.Cmd {
			case pb.CommandType_APPEND_LOG:
				ret.Msg += resp.Msg + "\n"
			default:
				log.Fatalln("no found cmd")
			}
		default:
			return ret
		}
	}
}
