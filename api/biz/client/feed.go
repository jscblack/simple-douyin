package client

import (
	"context"
	bizFeed "simple-douyin/api/biz/model/feed"
	"simple-douyin/api/biz/pack"
	"simple-douyin/kitex_gen/feed"
	"simple-douyin/kitex_gen/feed/feedservice"
	"simple-douyin/pkg/constant"
	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	apiLog "github.com/prometheus/common/log"
)

var feedClient feedservice.Client // interface from RPC IDL

func InitFeedClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := feedservice.NewClient(
		constant.FeedServiceName,
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
	feedClient = c
	apiLog.Info("Feed client initialized")
}

func Feed(ctx context.Context, req *feed.FeedRequest) (*bizFeed.FeedResponse, error) {
	resp, err := feedClient.Feed(ctx, req)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return pack.FeedPack(ctx, resp)
}
