package pack

import (
	"context"
	bizCommon "simple-douyin/api/biz/model/common"
	"simple-douyin/kitex_gen/common"

	apiLog "github.com/sirupsen/logrus"
)

func ToVideo(ctx context.Context, video *common.Video) *bizCommon.Video {
	if video == nil {
		return nil
	}

	return &bizCommon.Video{
		ID:            video.Id,
		Author:        UserPack(ctx, video.Author),
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

func ToVideos(ctx context.Context, videos []*common.Video) []*bizCommon.Video {
	videoList := make([]*bizCommon.Video, 0)
	for _, video := range videos {
		if v := ToVideo(ctx, video); v != nil {
			videoList = append(videoList, v)
		} else {
			apiLog.Fatal("video is nil!")
		}
	}
	return videoList
}
