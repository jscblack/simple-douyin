package dal

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/pkg/constant"
)

func QueryVideoFromVideoId(ctx context.Context, videoId int64) (*Video, error) {
	var video *Video
	if videoId < 0 {
		servLog.Error("videoId < 0!")
		return video, nil
	}
	if err := DB.Find(&video).Where("id = ?", videoId).Error; err != nil {
		servLog.Error("QueryVideoFromVideoId err", err)
		return video, err
	}
	servLog.Info("QueryVideoFromVideoId success")
	return video, nil
}

func QueryVideoFromUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var videoList []*Video
	if userId < 0 {
		servLog.Error("userId < 0!")
		return videoList, nil
	}
	if err := DB.Find(&videoList).Where("user_id = ?", userId).Order("id desc").Limit(constant.MaxListNum).Error; err != nil {
		servLog.Error("QueryVideoFromUserId err", err)
		return videoList, err
	}
	servLog.Info("QueryVideoFromUserId success")
	return videoList, nil
}
