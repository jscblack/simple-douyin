package main

import (
	"context"
	apiLog "github.com/prometheus/common/log"
	feed "simple-douyin/kitex_gen/feed"
	"simple-douyin/service/feed/service"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// TODO: Your code here...

	// 返回时间小于latestTime的最多30个视频
	// 可以根据req.userId推荐个性化视频
	if req.LatestTime == nil {
		timeStamp := time.Now().Unix()
		req.LatestTime = &timeStamp
	}

	resp, err := service.Feed(ctx, req)
	if err != nil {
		apiLog.Fatal(err)
	}

	return resp, nil
}
