package main

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/service"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp, err = service.PublishAction(ctx, req)
	if err != nil {
		servLog.Info(err)
		if resp == nil {
			resp = publish.NewPublishActionResponse()
		}
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return
	}
	servLog.Info("Video Published Successfully.")
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "Video Published Successfully."
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = service.PublishList(ctx, req)
	if err != nil {
		servLog.Info(err)
		if resp == nil {
			resp = publish.NewPublishListResponse()
		}
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "Get Video List Success."
	return resp, nil
}

// PublishVideoInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideoInfo(ctx context.Context, req *publish.PublishVideoInfoRequest) (resp *publish.PublishVideoInfoResponse, err error) {
	resp, err = service.PublishVideoInfo(ctx, req)
	if err != nil {
		servLog.Info(err)
		if resp == nil {
			resp = publish.NewPublishVideoInfoResponse()
		}
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return
	}
	servLog.Info("Get Video Info Success.")
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "Get Video Info Success."
	return
}

// PublishWorkCount implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishWorkCount(ctx context.Context, req *publish.PublishWorkCountRequest) (resp *publish.PublishWorkCountResponse, err error) {
	resp, err = service.PublishWorkCount(ctx, req)
	if err != nil {
		servLog.Info(err)
		if resp == nil {
			resp = publish.NewPublishWorkCountResponse()
		}
		resp.StatusCode = 57003
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return
	}
	servLog.Info("Get Work Count Success.")
	resp.StatusCode = 0
	resp.StatusMsg = new(string)
	*resp.StatusMsg = "Get Work Count Success."
	return
}
