package client

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
	"github.com/upyun/go-sdk/v3/upyun"
	"simple-douyin/kitex_gen/comment/commentservice"
	"simple-douyin/kitex_gen/favorite/favoriteservice"
	"simple-douyin/kitex_gen/user/userservice"
	"simple-douyin/pkg/constant"
	"time"
)

var UserClient userservice.Client // interface from RPC IDL
var FavoriteClient favoriteservice.Client
var CommentClient commentservice.Client
var UpyClient *upyun.UpYun

// Init 初始化，创建rpc client
func Init(ctx context.Context) {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}
	c1, err := userservice.NewClient(
		constant.UserServiceName,
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
	UserClient = c1
	servLog.Info("User client initialized")

	c2, err := favoriteservice.NewClient(
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
	FavoriteClient = c2
	servLog.Info("Favorite client initialized")

	c3, err := commentservice.NewClient(
		constant.CommentServiceName,
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
	CommentClient = c3
	servLog.Info("Comment client initialized")

	UpyClient = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   constant.Bucket,
		Operator: constant.Operator,
		Password: constant.UpyPassword,
	})
	servLog.Info("Upy client initialized")
}
