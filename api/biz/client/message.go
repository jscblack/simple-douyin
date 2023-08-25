package client

import (
	"simple-douyin/kitex_gen/message/messageservice"

	"context"
	bizMessage "simple-douyin/api/biz/model/message"
	"simple-douyin/api/biz/pack"
	kitexMessage "simple-douyin/kitex_gen/message"
	"simple-douyin/pkg/constant"
	"time"

	"github.com/cloudwego/kitex/client"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	etcd "github.com/kitex-contrib/registry-etcd"
	apiLog "github.com/sirupsen/logrus"
)

var messageClient messageservice.Client

func InitMessaggeClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := messageservice.NewClient(
		constant.MessageServiceName,
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		client.WithTracer(
			prometheus.NewClientTracer(
				constant.MessageClientTracerPort,
				constant.MessageClientTracerPath)), // tracer
		client.WithResolver(r), // resolver

	)
	if err != nil {
		apiLog.Fatal(err)
	}
	messageClient = c
	apiLog.Info("Message client initialized")
}

func MessageChat(ctx context.Context, bizReq *bizMessage.MessageChatRequest, bizResp *bizMessage.MessageChatResponse) error {
	rpcReq := &kitexMessage.MessageChatRequest{}
	err := pack.MessageChatUnpack(ctx, bizReq, rpcReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	rpcResp, err := messageClient.MessageChat(ctx, rpcReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.MessageChatPack(ctx, rpcResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func MessageSend(ctx context.Context, bizReq *bizMessage.MessageActionRequest, bizResp *bizMessage.MessageActionResponse) error {
	if bizReq.ActionType == 1 {
		rpcReq := &kitexMessage.MessageSendRequest{}
		err := pack.MessageSendUnpack(ctx, bizReq, rpcReq)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		rpcResp, err := messageClient.MessageSend(ctx, rpcReq)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		err = pack.MessageSendPack(ctx, rpcResp, bizResp)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		return nil
	} else {
		return nil
	}
}
