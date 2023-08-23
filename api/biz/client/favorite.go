package client

import (
	"context"
	bizFav "simple-douyin/api/biz/model/favorite"
	"simple-douyin/api/biz/pack"
	kiteFav "simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/favorite/favoriteservice"
	"simple-douyin/pkg/constant"
	"time"

	apiLog "github.com/prometheus/common/log"

	"github.com/cloudwego/kitex/client"

	etcd "github.com/kitex-contrib/registry-etcd"
)

var favoriteClient favoriteservice.Client // interface from RPC IDL

func InitFavoriteClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := favoriteservice.NewClient(
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
		apiLog.Fatal(err)
	}
	favoriteClient = c
	apiLog.Info("Favorite client initialized")
}

func FavoriteAdd(ctx context.Context, bizReq *bizFav.FavoriteActionRequest, bizResp *bizFav.FavoriteActionResponse) error {
	var err error
	kiteReq := new(kiteFav.FavoriteAddActionRequest)
	err = pack.FavoriteAddUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := favoriteClient.FavoriteAddAction(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.FavoriteAddpack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func FavoriteDel(ctx context.Context, bizReq *bizFav.FavoriteActionRequest, bizResp *bizFav.FavoriteActionResponse) error {
	var err error
	kiteReq := new(kiteFav.FavoriteDelActionRequest)
	err = pack.FavoriteDelUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := favoriteClient.FavoriteDelAction(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.FavoriteDelpack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func FavoriteList(ctx context.Context, bizReq *bizFav.FavoriteListRequest, bizResp *bizFav.FavoriteListResponse) error {
	var err error
	kiteReq := new(kiteFav.FavoriteListRequest)
	err = pack.FavoriteListUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := favoriteClient.FavoriteList(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.FavoriteListpack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}