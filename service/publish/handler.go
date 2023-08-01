package main

import (
	"context"
	"fmt"
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
		return nil, err
	}
	return
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp, err = service.PublishList(ctx, req)
	fmt.Println(resp.VideoList)
	if err != nil {
		servLog.Fatal(err)
		return nil, err
	}
	return resp, nil
}
