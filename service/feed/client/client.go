package client

import (
	"context"
	"simple-douyin/kitex_gen/comment/commentservice"
	"simple-douyin/kitex_gen/favorite/favoriteservice"
	"simple-douyin/kitex_gen/publish/publishservice"
	"simple-douyin/kitex_gen/user/userservice"
	"simple-douyin/pkg/constant"
	"time"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
)

var UserClient userservice.Client
var FavoriteClient favoriteservice.Client // interface from RPC IDL
var CommentClient commentservice.Client   // interface from RPC IDL
var PublishClient publishservice.Client

// 初始化，创建rpc client
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
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.UserClientTracerPort,
		// 		constant.UserClientTracerPath)), // tracer
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
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.FavoriteClientTracerPort,
		// 		constant.FavoriteClientTracerPath)), // tracer
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
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.CommentClientTracerPort,
		// 		constant.CommentClientTracerPath)), // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		servLog.Fatal(err)
	}
	CommentClient = c3
	servLog.Info("Comment client initialized")

	c4, err := publishservice.NewClient(
		constant.PublishServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithTracer(
		// 	prometheus.NewClientTracer(
		// 		constant.PublishClientTracerPort,
		// 		constant.PublishClientTracerPath)), // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		servLog.Fatal(err)
	}
	PublishClient = c4
	servLog.Info("Publish client initialized")
}
