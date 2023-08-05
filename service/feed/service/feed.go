package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/feed"
	"simple-douyin/service/feed/dal"
)

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// query from db according to req.LatestTime.

	videoList := make([]*common.Video, 0)

	dbVideoList, err := dal.QueryVideoFromLatestTime(ctx, req.GetLatestTime())
	if err != nil {
		servLog.Error("QueryVideoFromLatestTime err", err)
		return nil, err
	}

	for _, dbVideo := range dbVideoList {
		var video = &common.Video{
			Id: int64(dbVideo.ID),
			// Author:        get from UserInfo()
			PlayUrl:  dbVideo.PlayUrl,
			CoverUrl: dbVideo.CoverUrl,
			// FavoriteCount: get from VideoFavorCount()
			// CommentCount:  get from CommentCount()
			// IsFavorite:    get from IsFavor()
			Title: dbVideo.Title,
		}
		videoList = append(videoList, video)
	}

	// for test
	//author := common.User{
	//	Id:   1,
	//	Name: "Koschei",
	//	// FollowCount:   10022,
	//	// FollowerCount: 3,
	//	IsFollow: true,
	//}
	//video := common.Video{
	//	Id:            1,
	//	Author:        &author,
	//	PlayUrl:       "http://a.com",
	//	CoverUrl:      "http://b.com",
	//	FavoriteCount: 200,
	//	CommentCount:  200,
	//	IsFavorite:    true,
	//	Title:         "bear",
	//}
	//
	//videoList = append(videoList, &video)

	// 取返回视频中最早的时间作为next time
	earliestTime := dbVideoList[len(dbVideoList)-1].CreateTime
	return &feed.FeedResponse{
		StatusCode: 0,
		VideoList:  videoList,
		NextTime:   &earliestTime,
	}, nil
}
