package dal

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/pkg/constant"
	"time"
)

func QueryVideoFromLatestTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	var videoList []*Video
	if latestTime < 0 {
		servLog.Error("latestTime < 0!")
		return videoList, nil
	}
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	if err := DB.Find(&videoList).Where("create_time < ?", latestTime).Order("id desc").Limit(constant.MaxFeedNum).Error; err != nil {
		servLog.Error("QueryVideoFromLatestTime err", err)
		return videoList, err
	}
	servLog.Info("QueryVideoFromLatestTime success")
	return videoList, nil
}
