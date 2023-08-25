package client

import (
	"context"
	bizPublish "simple-douyin/api/biz/model/publish"
	"simple-douyin/api/biz/pack"
	kitexPublish "simple-douyin/kitex_gen/publish"
	"simple-douyin/kitex_gen/publish/publishservice"
	"simple-douyin/pkg/constant"
	"time"

	apiLog "github.com/sirupsen/logrus"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"

	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client // interface from RPC IDL

func InitPublishClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := publishservice.NewClient(
		constant.PublishServiceName,
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
	publishClient = c
	apiLog.Info("Publish client initialized")
}

func PublishAction(ctx context.Context, req *kitexPublish.PublishActionRequest) (*bizPublish.PublishActionResponse, error) {
	resp, err := publishClient.PublishAction(ctx, req, callopt.WithRPCTimeout(time.Minute))
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return pack.PublishActionPack(ctx, resp)
}

func PublishList(ctx context.Context, req *kitexPublish.PublishListRequest) (*bizPublish.PublishListResponse, error) {
	resp, err := publishClient.PublishList(ctx, req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}
	return pack.PublishListPack(ctx, resp)
}
