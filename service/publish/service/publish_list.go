package service

import (
	"context"
	"simple-douyin/kitex_gen/common"
	publish "simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/dal"
	"sync"

	servLog "github.com/sirupsen/logrus"
)

// PublishList implements the PublishServiceImpl interface.
func PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// query db for videoList according to userId.
	servLog.Info("Accept request: ", req)

	dbVideoList, err := dal.QueryVideoListFromUserId(ctx, req.UserId)
	servLog.Info("list:", dbVideoList)
	if err != nil {
		servLog.Error("QueryVideoFromUserId err", err)
		return nil, err
	}

	// for _, dbVideo := range dbVideoList {
	// 	video, err := fillVideoInfo(ctx, dbVideo, &req.FromUserId)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	videoList = append(videoList, video)
	// }
	videoList := make([]*common.Video, len(dbVideoList))
	var mu sync.Mutex
	var wg sync.WaitGroup
	for idx, dbVideo := range dbVideoList {
		wg.Add(1)
		go func(idx int, dbVideo *dal.Video) {
			defer wg.Done()
			video, err := fillVideoInfo(ctx, dbVideo, &req.FromUserId)
			if err != nil {
				servLog.Error("fillVideoInfo err", err)
				return
			}
			mu.Lock()
			videoList[idx] = video
			mu.Unlock()

		}(idx, dbVideo)
	}
	wg.Wait() // Wait for all goroutines to complete

	resp = publish.NewPublishListResponse()
	resp.VideoList = videoList
	servLog.Info("PublishList success.")
	return resp, nil
}
