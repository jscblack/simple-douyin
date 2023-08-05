package client

import (
	"context"
	"simple-douyin/kitex_gen/favorite/favoriteservice"
	"simple-douyin/kitex_gen/relation/relationservice"
	"simple-douyin/pkg/constant"
	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/prometheus/common/log"
)

var FavoriteClient favoriteservice.Client // interface from RPC IDL
var RelationClient relationservice.Client // interface from RPC IDL
// 初始化，创建rpc client
func Init(ctx context.Context) {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}
	c1, err := favoriteservice.NewClient(
		constant.FavoriteServiceName,
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
		servLog.Fatal(err)
	}
	FavoriteClient = c1
	servLog.Info("Favorite client initialized")

	c2, err := relationservice.NewClient(
		constant.RelationServiceName,
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
		servLog.Fatal(err)
	}
	RelationClient = c2
	servLog.Info("Relation client initialized")
}
