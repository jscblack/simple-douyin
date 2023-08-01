package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	publish "simple-douyin/kitex_gen/publish"
)

func PublishVideoInfo(ctx context.Context, req *publish.PublishVideoInfoRequest) (resp *publish.PublishVideoInfoResponse, err error) {
	// query video info from db according to videoId.
	servLog.Info("Accept request: ", req)
	return
}
