package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/common"
	publish "simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/dal"
)

// PublishList implements the PublishServiceImpl interface.
func PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// query db for videoList according to userId.
	// req.UserId
	servLog.Info("Accept request: ", req)

	videoList := make([]*common.Video, 0)
	dbVideoList, err := dal.QueryVideoFromUserId(ctx, req.UserId)
	if err != nil {
		servLog.Error("QueryVideoFromUserId err", err)
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
	//videoList = append(videoList, &video)

	return &publish.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  videoList,
	}, nil
}
