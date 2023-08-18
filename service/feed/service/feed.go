package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/comment"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/feed"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/feed/client"
	"simple-douyin/service/feed/dal"
)

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// query from db according to req.LatestTime.
	servLog.Info("Feed rpc.")

	var videoList []*common.Video
	dbVideoList, err := dal.QueryVideoFromLatestTime(ctx, req.GetLatestTime())
	servLog.Info("after query.")
	if err != nil {
		servLog.Error("QueryVideoFromLatestTime err", err)
		return nil, err
	}

	for _, dbVideo := range dbVideoList {
		video, err := fillVideoInfo(ctx, dbVideo)
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}

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

func fillVideoInfo(ctx context.Context, dbVideo *dal.Video) (*common.Video, error) {
	servLog.Info("Rpc userInfo.")
	userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{ToUserId: dbVideo.UserId})
	servLog.Info(userResp)
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
	isFavResp, err := client.FavoriteClient.IsFavor(ctx, &favorite.IsFavorRequest{UserId: dbVideo.UserId, VideoId: dbVideo.ID})
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
