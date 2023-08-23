package main

import (
	"context"
	feed "simple-douyin/kitex_gen/feed"
	"simple-douyin/service/feed/service"

	apiLog "github.com/prometheus/common/log"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// 返回时间小于latestTime的最多30个视频
	// 可以根据req.userId推荐个性化视频

	resp, err := service.Feed(ctx, req)
	if err != nil {
		apiLog.Error(err)
		return nil, err
	}

	return resp, nil
}
