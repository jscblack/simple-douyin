package service

import (
	"context"
	"simple-douyin/kitex_gen/comment"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/feed"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/feed/client"
	"simple-douyin/service/feed/dal"
	"sync"

	servLog "github.com/sirupsen/logrus"
)

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// query from db according to req.LatestTime.
	servLog.Info("Feed rpc.")

	dbVideoList, err := dal.QueryVideoFromLatestTime(ctx, req.GetLatestTime())
	servLog.Info("after query.")
	if err != nil {
		servLog.Error("QueryVideoFromLatestTime err", err)
		return nil, err
	}

	// for _, dbVideo := range dbVideoList {
	// 	video, err := fillVideoInfo(ctx, dbVideo, req.UserId)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	videoList = append(videoList, video)
	// }
	// use go routine to fill each video info concurrently
	// and use channel to receive the result
	// keep the order of videoList
	videoList := make([]*common.Video, len(dbVideoList))
	var mu sync.Mutex
	var wg sync.WaitGroup
	for idx, dbVideo := range dbVideoList {
		wg.Add(1)
		go func(idx int, dbVideo *dal.Video) {
			defer wg.Done()
			video, err := fillVideoInfo(ctx, dbVideo, req.UserId)
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

	servLog.Info("success info.")
	// 取返回视频中最早的时间作为next time
	var earliestTime int64
	if len(videoList) > 0 {
		earliestTime, err = dal.QueryEarliestTimeFromVideoList(ctx, videoList)
		if err != nil {
			return nil, err
		}
	}
	servLog.Info("Feed Success.")
	return &feed.FeedResponse{
		StatusCode: 0,
		VideoList:  videoList,
		NextTime:   &earliestTime,
	}, nil
}

func fillVideoInfo(ctx context.Context, dbVideo *dal.Video, userId *int64) (*common.Video, error) {
	servLog.Info("Rpc userInfo.")
	userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
		UserId:   userId,
		ToUserId: dbVideo.UserId,
	})
	servLog.Info(userResp)
	servLog.Info(userResp.User.WorkCount)
	if err != nil {
		return nil, err
	}

	servLog.Info("Rpc favored_count.")
	favResp, err := client.FavoriteClient.VideoFavoredCount(ctx, &favorite.VideoFavoredCountRequest{VideoId: dbVideo.ID})
	if err != nil {
		return nil, err
	}

	servLog.Info("Rpc comment_count.")
	comResp, err := client.CommentClient.CommentCount(ctx, &comment.CommentCountRequest{VideoId: dbVideo.ID})
	if err != nil {
		return nil, err
	}

	servLog.Info("Rpc is_favored.")
	isFavResp, err := client.FavoriteClient.IsFavor(ctx, &favorite.IsFavorRequest{UserId: *userId, VideoId: dbVideo.ID})
	if err != nil {
		return nil, err
	}

	return &common.Video{
		Id:            dbVideo.ID,
		Author:        userResp.User,
		PlayUrl:       dbVideo.PlayUrl,
		CoverUrl:      dbVideo.CoverUrl,
		FavoriteCount: favResp.GetFavoredCount(),
		CommentCount:  comResp.GetCommentCount(),
		IsFavorite:    isFavResp.IsFavorite,
		Title:         dbVideo.Title,
	}, nil
}
