package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	publish "simple-douyin/kitex_gen/publish"
)

func PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// 根据userId向db写入视频的data和title
	servLog.Info("Accept request: ", req)

	return &publish.PublishActionResponse{
		StatusCode: 200,
		StatusMsg:  nil,
	}, nil
}
