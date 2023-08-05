package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/dal"
)

func PublishVideoInfo(ctx context.Context, req *publish.PublishVideoInfoRequest) (resp *publish.PublishVideoInfoResponse, err error) {
	// query video info from db according to videoId.
	servLog.Info("Accept request: ", req)

	dbVideo, err := dal.QueryVideoFromVideoId(ctx, req.VideoId)
	if err != nil {
		servLog.Error("QueryVideoFromVideoId err", err)
		return
	}

	// for test
	author := common.User{
		Id:   1,
		Name: "Koschei",
		// FollowCount:   10022, // get from RelationFollowCount()
		// FollowerCount: 3,     // get from RelationFollowerCount()
		IsFollow: true,
	}

	resp.Video = &common.Video{
		Id:            int64(dbVideo.ID),
		Author:        &author, // get from UserInfo()
		PlayUrl:       dbVideo.PlayUrl,
		CoverUrl:      dbVideo.CoverUrl,
		FavoriteCount: 0,    // get from VideoFavorCount()
		CommentCount:  0,    // get from CommentCount()
		IsFavorite:    true, // get from IsFavor()
		Title:         dbVideo.Title,
	}
	return
}
