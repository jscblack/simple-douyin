package service

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/comment"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/publish/client"
	"simple-douyin/service/publish/dal"
)

func PublishVideoInfo(ctx context.Context, req *publish.PublishVideoInfoRequest) (resp *publish.PublishVideoInfoResponse, err error) {
	// query video info from db according to videoId.
	servLog.Info("Accept request: ", req)

	var dbVideo *dal.Video
	dbVideo, err = dal.QueryVideoFromVideoId(ctx, req.VideoId)
	if err != nil {
		servLog.Error("QueryVideoFromVideoId err", err)
		return nil, err
	}
	resp = publish.NewPublishVideoInfoResponse()
	resp.Video, err = fillVideoInfo(ctx, dbVideo)
	if err != nil {
		return nil, err
	}
	return
}

func fillVideoInfo(ctx context.Context, dbVideo *dal.Video) (*common.Video, error) {
	servLog.Info("Rpc userInfo.")
	userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{ToUserId: dbVideo.UserId})
	if err != nil {
		return nil, err
	}

	servLog.Info("Rpc favored_count.")
	favResp, err := client.FavoriteClient.VideoFavoredCount(ctx, &favorite.VideoFavoredCountRequest{dbVideo.ID})
	if err != nil {
		return nil, err
	}

	servLog.Info("Rpc comment_count.")
	comResp, err := client.CommentClient.CommentCount(ctx, &comment.CommentCountRequest{dbVideo.ID})
	if err != nil {
		return nil, err
	}

	servLog.Info("Rpc is_favored.")
	isFavResp, err := client.FavoriteClient.IsFavor(ctx, &favorite.IsFavorRequest{dbVideo.UserId, dbVideo.ID})
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
