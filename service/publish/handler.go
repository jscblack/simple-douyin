package main

import (
	"context"
	servLog "github.com/prometheus/common/log"
	publish "simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/service"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp, err = service.PublishAction(ctx, req)
	if err != nil {
		servLog.Fatal(err)
	}
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = service.PublishList(ctx, req)
	if err != nil {
		servLog.Fatal(err)
	}
	return
}

// PublishVideoInfo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideoInfo(ctx context.Context, req *publish.PublishVideoInfoRequest) (resp *publish.PublishVideoInfoResponse, err error) {
	resp, err = service.PublishVideoInfo(ctx, req)
	if err != nil {
		servLog.Fatal(err)
	}
	return
}
