package main

import (
	"context"
	comment "simple-douyin/kitex_gen/comment"
	"simple-douyin/service/comment/service"

	servLog "github.com/sirupsen/logrus"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAddAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAddAction(ctx context.Context, req *comment.CommentAddActionRequest) (resp *comment.CommentAddActionResponse, err error) {
	resp = new(comment.CommentAddActionResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.CommentAdd(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57005
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

// CommentDelAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentDelAction(ctx context.Context, req *comment.CommentDelActionRequest) (resp *comment.CommentDelActionResponse, err error) {
	resp = new(comment.CommentDelActionResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.CommentDel(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57005
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

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = new(comment.CommentListResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.CommentList(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57005
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

// CommentCount implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentCount(ctx context.Context, req *comment.CommentCountRequest) (resp *comment.CommentCountResponse, err error) {
	resp = new(comment.CommentCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.CommentCount(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57005
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
	}
	// 后处理返回结果
	// ...
	return resp, nil
}
