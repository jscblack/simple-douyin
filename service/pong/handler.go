package main

import (
	"context"
	pong "simple-douyin/kitex_gen/pong"
	"time"
)

// PongServiceImpl implements the last service interface defined in the IDL.
type PongServiceImpl struct{}

// Pong implements the PongServiceImpl interface.
func (s *PongServiceImpl) Pong(ctx context.Context, req *pong.PingReq) (resp *pong.PongResp, err error) {
	// TODO: Your code here...
	resp = &pong.PongResp{
		PongTime: req.PingTime + " " + time.Now().String(),
	}
	return resp, nil
}
