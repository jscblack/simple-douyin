package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	publish "simple-douyin/kitex_gen/publish"
)

func PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	// 根据userId向db写入视频数据
	// 需要生成PlayUrl和CoverUrl
	servLog.Info("Accept request: ", req)

	return &publish.PublishActionResponse{
		StatusCode: 0,
		StatusMsg:  nil,
	}, nil
}
