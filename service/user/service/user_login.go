package service

import (
	"context"
	"simple-douyin/kitex_gen/user"

	servLog "github.com/prometheus/common/log"
)

func UserLogin(ctx context.Context, req *user.UserLoginRequest, resp *user.UserLoginResponse) error {
	// 实际业务
	servLog.Info("Accept request: ", req)
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.UserId = 114514
	return nil
}
