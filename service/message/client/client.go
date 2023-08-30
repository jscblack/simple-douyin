package client

import (
	"context"
	"simple-douyin/kitex_gen/relation/relationservice"
	"simple-douyin/pkg/constant"
	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
)

var RelationClient relationservice.Client // interface from RPC IDL
// 初始化，创建rpc client
func Init(ctx context.Context) {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}
	c1, err := relationservice.NewClient(
		constant.RelationServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.RelationClientTracerPort,
		// 		constant.RelationClientTracerPath)), // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		servLog.Fatal(err)
	}
	RelationClient = c1
	servLog.Info("Relation client initialized")
}
