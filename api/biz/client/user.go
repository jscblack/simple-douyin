package client

import (
	"context"
	bizUser "simple-douyin/api/biz/model/user"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/kitex_gen/user/userservice"
	"simple-douyin/pkg/constant"
	"time"

	apiLog "github.com/prometheus/common/log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"

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
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		apiLog.Fatal(err)
	}
	userClient = c
	apiLog.Info("User client initialized")
}

func UserRegister(ctx context.Context, bizReq *bizUser.UserRegisterRequest, bizResp *bizUser.UserRegisterResponse) error {
	req := user.UserRegisterRequest{
		Username: bizReq.Username,
		Password: bizReq.Password,
	}
	resp, err := userClient.UserRegister(ctx, &req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		apiLog.Error(err)
		return err
	}
	bizResp.StatusCode = resp.StatusCode
	bizResp.StatusMsg = resp.StatusMsg
	bizResp.UserID = resp.UserId
	return nil
}

func UserLogin(ctx context.Context, bizReq *bizUser.UserLoginRequest, bizResp *bizUser.UserLoginResponse) error {
	req := user.UserLoginRequest{
		Username: bizReq.Username,
		Password: bizReq.Password,
	}
	resp, err := userClient.UserLogin(ctx, &req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		apiLog.Error(err)
		return err
	}
	bizResp.StatusCode = resp.StatusCode
	bizResp.StatusMsg = resp.StatusMsg
	bizResp.UserID = resp.UserId
	return nil
}
