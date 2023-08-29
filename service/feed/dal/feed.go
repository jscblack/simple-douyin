package dal

import (
	"context"
	"encoding/json"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/pkg/constant"
	"strconv"
	"time"

	servLog "github.com/sirupsen/logrus"
)

func QueryVideoFromLatestTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	servLog.Info("Query begin.")
	var videoList []*Video
	if latestTime < 0 {
		servLog.Error("latestTime < 0!")
		return videoList, nil
	}
	if latestTime == 0 {
		latestTime = time.Now().UnixMilli()
	}
	servLog.Info("latestTime: ", time.UnixMilli(latestTime))
	if err := DB.Where("created_at < ?", time.UnixMilli(latestTime)).Order("id desc").Limit(constant.MaxFeedNum).Find(&videoList).Error; err != nil {
		servLog.Error("QueryVideoFromLatestTime err", err)
		return videoList, err
	}
	servLog.Info("QueryVideoFromLatestTime success")
	return videoList, nil
}

func QueryEarliestTimeFromVideoList(ctx context.Context, videoList []*common.Video) (int64, error) {
	var video *Video
	servLog.Info("QueryEarliestTimeFromVideoList begin.")
	videoId := videoList[len(videoList)-1].Id
	servLog.Info("video id: ", videoId)

	videoKey := "v" + strconv.FormatInt(videoId, 10)
	cacheVideo, err := RDB.Get(ctx, videoKey).Result()
	servLog.Info(cacheVideo)
	servLog.Info(err)
	if err != nil {
		// 缓存不存在
		servLog.Warn("Video Info Not Found In Cache: ", videoId)
		if err = DB.Where("id = ?", videoId).Take(&video).Error; err != nil {
			servLog.Error("QueryEarliestTimeFromVideoList err", err)
			return 0, err
		}
		servLog.Info("QueryEarliestTimeFromVideoList success")

		// 写入redis缓存
		videoJson, err := json.Marshal(video)
		if err != nil {
			return 0, err
		}
		err = RDB.Set(ctx, videoKey, videoJson, 0).Err()
		if err != nil {
			servLog.Error(err)
			return 0, err
		}
		return video.CreatedAt.UnixMilli(), nil
	}
	err = json.Unmarshal([]byte(cacheVideo), &video)
	if err != nil {
		return 0, err
	}
	servLog.Info("QueryEarliestTimeFromVideoList success.")
	return video.CreatedAt.UnixMilli(), nil
}
