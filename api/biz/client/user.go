package client

import (
	"context"
	bizUser "simple-douyin/api/biz/model/user"
	"simple-douyin/api/biz/pack"
	kiteUser "simple-douyin/kitex_gen/user"
	"simple-douyin/kitex_gen/user/userservice"
	"simple-douyin/pkg/constant"
	"time"

	apiLog "github.com/sirupsen/logrus"

	"github.com/cloudwego/kitex/client"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client // interface from RPC IDL

func InitUserClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := userservice.NewClient(
		constant.UserServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithTracer(
			prometheus.NewClientTracer(
				constant.UserClientTracerPort,
				constant.UserClientTracerPath)), // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		apiLog.Fatal(err)
	}
	userClient = c
	apiLog.Info("User client initialized")
}

func UserRegister(ctx context.Context, bizReq *bizUser.UserRegisterRequest, bizResp *bizUser.UserRegisterResponse) error {
	var err error
	kiteReq := new(kiteUser.UserRegisterRequest)
	err = pack.UserRegisterUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := userClient.UserRegister(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.UserRegisterPack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func UserLogin(ctx context.Context, bizReq *bizUser.UserLoginRequest, bizResp *bizUser.UserLoginResponse) error {
	var err error
	kiteReq := new(kiteUser.UserLoginRequest)
	err = pack.UserLoginUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := userClient.UserLogin(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.UserLoginPack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}

func UserInfo(ctx context.Context, bizReq *bizUser.UserInfoRequest, bizResp *bizUser.UserInfoResponse) error {
	var err error
	kiteReq := new(kiteUser.UserInfoRequest)
	err = pack.UserInfoUnpack(ctx, bizReq, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	kiteResp, err := userClient.UserInfo(ctx, kiteReq)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	err = pack.UserInfoPack(ctx, kiteResp, bizResp)
	if err != nil {
		apiLog.Error(err)
		return err
	}
	return nil
}
