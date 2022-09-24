package server

import (
	pb "where_to_eat/network/protobuf"
)

func EventExecDispatcher(in chan *pb.Req, m map[int64]chan *pb.Resp) {
	event := <-in
	if _, ok := m[event.ClientId]; !ok {
		m[event.ClientId] = make(chan *pb.Resp, 16)
	}
	go eventHandler(event, m)
}

func eventHandler(req *pb.Req, m map[int64]chan *pb.Resp) {
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
	select {
	case resp := <-respCh:
		ret.Msg += resp.Msg
	default:
	}
	return ret
}
