package pack

import (
	apiLog "github.com/prometheus/common/log"
	bizCommon "simple-douyin/api/biz/model/common"
	bizPublish "simple-douyin/api/biz/model/publish"
	"simple-douyin/kitex_gen/common"
	kitexPublish "simple-douyin/kitex_gen/publish"
)

func ToVideo(video *common.Video) *bizCommon.Video {
	if video == nil {
		return nil
	}

	author := video.Author
	user := &bizCommon.User{
		ID:            author.Id,
		Name:          author.Name,
		FollowCount:   author.FollowerCount,
		FollowerCount: author.FollowerCount,
		IsFollow:      author.IsFollow,
	}

	return &bizCommon.Video{
		ID:            video.Id,
		Author:        user,
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

func ToVideos(videos []*common.Video) []*bizCommon.Video {
	videoList := make([]*bizCommon.Video, 0)
	for _, video := range videos {
		if v := ToVideo(video); v != nil {
			videoList = append(videoList, v)
		} else {
			apiLog.Fatal("video is nil!")
		}
	}
	return videoList
}

// PublishActionPack kitexReq -> bizReq
func PublishActionPack(kitexReq *kitexPublish.PublishActionResponse) (*bizPublish.PublishActionResponse, error) {
	return &bizPublish.PublishActionResponse{
		StatusCode: kitexReq.StatusCode,
		StatusMsg:  kitexReq.StatusMsg,
	}, nil
}

// PublishListPack kitexReq -> bizReq
func PublishListPack(kitexReq *kitexPublish.PublishListResponse) (*bizPublish.PublishListResponse, error) {
	return &bizPublish.PublishListResponse{
		StatusCode: kitexReq.StatusCode,
		StatusMsg:  kitexReq.StatusMsg,
		VideoList:  ToVideos(kitexReq.VideoList),
	}, nil
}
