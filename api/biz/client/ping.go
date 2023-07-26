package client

import (
	"context"
	"log"
	"simple-douyin/kitex_gen/pong"
	"simple-douyin/kitex_gen/pong/pongservice"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func PingClient() string {
	c, err := pongservice.NewClient("example-server", client.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		log.Fatal(err)
	}
	req := &pong.PingReq{
		PingTime: time.Now().String(),
	}
	resp, err := c.Pong(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	return resp.PongTime
}
