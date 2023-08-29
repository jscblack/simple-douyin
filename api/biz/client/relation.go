package client

import (
	"context"
	bizRelation "simple-douyin/api/biz/model/relation"
	"simple-douyin/api/biz/pack"
	kiteRelation "simple-douyin/kitex_gen/relation"
	"simple-douyin/kitex_gen/relation/relationservice"
	"simple-douyin/pkg/constant"
	"time"

	"github.com/cloudwego/kitex/client"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	etcd "github.com/kitex-contrib/registry-etcd"
	apiLog "github.com/sirupsen/logrus"
)

var relationClient relationservice.Client

func InitRelationClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := relationservice.NewClient(
		constant.RelationServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithTracer(
			prometheus.NewClientTracer(
				constant.RelationClientTracerPort,
				constant.RelationClientTracerPath)),
		client.WithResolver(r),
	)
	if err != nil {
		apiLog.Fatal(err)
	}
	relationClient = c
	apiLog.Info("Relation client initialized")
}

func RelationAction(ctx context.Context, bizReq *bizRelation.RelationActionRequest, bizResp *bizRelation.RelationActionResponse) error {
	var err error
	reqType := bizReq.ActionType
	if reqType == 1 {
		kiteReq := new(kiteRelation.RelationAddRequest)
		err = pack.RelationAddUnpack(ctx, bizReq, kiteReq)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		if relationClient == nil {
			apiLog.Error("client/relation :NIL !!!!")
			return err
		}
		kiteResp, err := relationClient.RelationAdd(ctx, kiteReq)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		err = pack.RelationAddPack(ctx, kiteResp, bizResp)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		return nil
	} else {
		kiteReq := new(kiteRelation.RelationRemoveRequest)
		err = pack.RelationRemoveUnpack(ctx, bizReq, kiteReq)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		kiteResp, err := relationClient.RelationRemove(ctx, kiteReq)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		err = pack.RelationRemovePack(ctx, kiteResp, bizResp)
		if err != nil {
			apiLog.Error(err)
			return err
		}
		return nil
	}
}

func RelationFollowList(ctx context.Context, bizReq *bizRelation.RelationFollowListRequest, bizResp *bizRelation.RelationFollowListResponse) error {
	var err error
	kiteReq := new(kiteRelation.RelationFollowListRequest)
	err = pack.RelationFollowListUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := relationClient.RelationFollowList(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.RelationFollowListPack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func RelationFollowerList(ctx context.Context, bizReq *bizRelation.RelationFollowerListRequest, bizResp *bizRelation.RelationFollowerListResponse) error {
	var err error
	kiteReq := new(kiteRelation.RelationFollowerListRequest)
	err = pack.RelationFollowerListUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := relationClient.RelationFollowerList(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.RelationFollowerListPack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func RelationFriendList(ctx context.Context, bizReq *bizRelation.RelationFriendListRequest, bizResp *bizRelation.RelationFriendListResponse) error {
	var err error
	kiteReq := new(kiteRelation.RelationFriendListRequest)
	err = pack.RelationFriendListUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := relationClient.RelationFriendList(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.RelationFriendListPack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}
