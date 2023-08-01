package main

import (
	"context"
	user "simple-douyin/kitex_gen/user"
	"simple-douyin/service/user/service"

	servLog "github.com/prometheus/common/log"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)

	// 前处理校验请求
	if len(req.Username) == 0 {
		resp.StatusCode = 57001 // status code 用来描述业务错误，5+服务端口
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "用户名不能为空"
		return resp, nil
	}
	if len(req.Password) == 0 {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "密码不能为空"
		return resp, nil
	}
	if len(req.Username) > 32 {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "用户名不能超过32位"
		return resp, nil
	}
	if len(req.Password) > 32 {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "密码不能超过32位"
		return resp, nil
	}

	// 实际业务
	err = service.UserRegister(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)

	// 前处理校验请求
	if len(req.Username) == 0 {
		resp.StatusCode = 57001 // status code 用来描述业务错误，5+服务端口
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "用户名不能为空"
		return resp, nil
	}
	if len(req.Password) == 0 {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "密码不能为空"
		return resp, nil
	}
	// 实际业务
	err = service.UserLogin(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	resp = new(user.UserInfoResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.UserInfo(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}

// UpdateUserCounter implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserCounter(ctx context.Context, req *user.UpdateUserCounterRequest) (resp *user.UpdateUserCounterResponse, err error) {
	resp = new(user.UpdateUserCounterResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.UpdateUserCounter(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}
