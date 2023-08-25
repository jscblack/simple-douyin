package client

import (
	"context"
	bizComment "simple-douyin/api/biz/model/comment"
	"simple-douyin/api/biz/pack"
	kiteComment "simple-douyin/kitex_gen/comment"
	"simple-douyin/kitex_gen/comment/commentservice"
	"simple-douyin/pkg/constant"
	"time"

	apiLog "github.com/sirupsen/logrus"

	"github.com/cloudwego/kitex/client"

	etcd "github.com/kitex-contrib/registry-etcd"
)

var commentClient commentservice.Client // interface from RPC IDL

func InitCommentClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := commentservice.NewClient(
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
		apiLog.Fatal(err)
	}
	commentClient = c
	apiLog.Info("Comment client initialized")
}

func CommentAdd(ctx context.Context, bizReq *bizComment.CommentActionRequest, bizResp *bizComment.CommentActionResponse) error {
	var err error
	kiteReq := new(kiteComment.CommentAddActionRequest)
	err = pack.CommentAddUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := commentClient.CommentAddAction(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.CommentAddpack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func CommentDel(ctx context.Context, bizReq *bizComment.CommentActionRequest, bizResp *bizComment.CommentActionResponse) error {
	var err error
	kiteReq := new(kiteComment.CommentDelActionRequest)
	err = pack.CommentDelUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := commentClient.CommentDelAction(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.CommentDelpack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func CommentList(ctx context.Context, bizReq *bizComment.CommentListRequest, bizResp *bizComment.CommentListResponse) error {
	var err error
	kiteReq := new(kiteComment.CommentListRequest)
	err = pack.CommentListUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := commentClient.CommentList(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.CommentListpack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}
