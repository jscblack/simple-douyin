package service

import (
	"context"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/feed"
	"time"
)

func Feed(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	// query from db.
	// req.LatestTime

	videoList := make([]*common.Video, 0)

	author := common.User{
		Id:   1,
		Name: "Koschei",
		// FollowCount:   10022,
		// FollowerCount: 3,
		IsFollow: true,
	}
	video := common.Video{
		Id:            1,
		Author:        &author,
		PlayUrl:       "http://a.com",
		CoverUrl:      "http://b.com",
		FavoriteCount: 200,
		CommentCount:  200,
		IsFavorite:    true,
		Title:         "bear",
	}

	videoList = append(videoList, &video)

	// 取返回视频中最早的时间作为next time

	earliestTime := time.Now().Unix()
	return &feed.FeedResponse{
		StatusCode: 200,
		VideoList:  videoList,
		NextTime:   &earliestTime,
	}, nil
}
