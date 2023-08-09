package dal

import (
	"context"
	"encoding/json"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/pkg/constant"
	"strconv"
	"time"
)

func QueryVideoFromLatestTime(ctx context.Context, latestTime int64) ([]*Video, error) {
	servLog.Info("Query begin.")
	var videoList []*Video
	if latestTime < 0 {
		servLog.Error("latestTime < 0!")
		return videoList, nil
	}
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	if err := DB.Where("created_at < ?", time.Unix(latestTime, 0)).Order("id desc").Limit(constant.MaxFeedNum).Find(&videoList).Error; err != nil {
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
	cacheVideo, err := RDB.Get(ctx, strconv.FormatInt(videoId, 10)).Result()
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
		err = RDB.Set(ctx, strconv.FormatInt(videoId, 10), videoJson, 0).Err()
		if err != nil {
			servLog.Error(err)
			return 0, err
		}
		return video.CreatedAt.Unix(), nil
	}
	err = json.Unmarshal([]byte(cacheVideo), &video)
	if err != nil {
		return 0, err
	}
	servLog.Info("QueryEarliestTimeFromVideoList success.")
	return video.CreatedAt.Unix(), nil
}
