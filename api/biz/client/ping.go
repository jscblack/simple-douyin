package client

import (
	"context"

	"simple-douyin/kitex_gen/pong"
	"simple-douyin/kitex_gen/pong/pongservice"
	"simple-douyin/pkg/constant"
	"time"

	apiLog "github.com/sirupsen/logrus"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"

	etcd "github.com/kitex-contrib/registry-etcd"
)

var pingClient pongservice.Client // interface from RPC IDL

func InitPingClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := pongservice.NewClient(
		constant.PingServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		apiLog.Fatal(err)
	}
	pingClient = c
	apiLog.Info("Ping client initialized")
}

func PingClient() string {
	req := &pong.PingReq{
		PingTime: time.Now().String(),
	}
	resp, err := pingClient.Pong(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		apiLog.Error(err)
	}
	return resp.PongTime
}
