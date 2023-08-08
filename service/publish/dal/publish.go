package dal

import (
	"context"
	"encoding/json"
	servLog "github.com/prometheus/common/log"
	"simple-douyin/pkg/constant"
	"strconv"
)

func QueryVideoFromVideoId(ctx context.Context, videoId int64) (*Video, error) {
	var video *Video
	if videoId < 0 {
		servLog.Error("videoId < 0!")
		return video, nil
	}
	// query from Redis.
	cacheVideo, err := RDB.Get(ctx, strconv.FormatInt(videoId, 10)).Result()
	if err != nil {
		// 缓存不存在
		servLog.Warn("Video Info Not Found In Cache: ", videoId)
		// query from db.
		if err := DB.Where("id = ?", videoId).Take(&video).Error; err != nil {
			servLog.Error("QueryVideoFromVideoId err", err)
			return nil, err
		}
		servLog.Info("QueryVideoFromVideoId success")

		// 写入redis缓存
		videoJson, err := json.Marshal(video)
		if err != nil {
			return nil, err
		}
		err = RDB.Set(ctx, strconv.FormatInt(videoId, 10), videoJson, 0).Err()
		if err != nil {
			servLog.Error(err)
			return nil, err
		}
		return video, nil
	}
	// 缓存存在
	servLog.Info("Video Get From Cache: ", cacheVideo)
	// maybe crashed.
	err = json.Unmarshal([]byte(cacheVideo), &video)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func QueryVideoListFromUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var videoList []*Video
	if userId < 0 {
		servLog.Error("userId < 0!")
		return videoList, nil
	}
	if err := DB.Where("user_id = ?", userId).Order("id desc").Limit(constant.MaxListNum).Find(&videoList).Error; err != nil {
		servLog.Error("QueryVideoFromUserId err", err)
		return videoList, err
	}
	servLog.Info("QueryVideoFromUserId success")
	return videoList, nil
}

func QueryWorkCountFromUserId(ctx context.Context, userId int64) (int64, error) {
	var workCount int64
	if userId < 0 {
		servLog.Error("userId < 0!")
		return workCount, nil
	}

	// query from Redis.
	cacheWorkCount, err := RDB.Get(ctx, strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		// 缓存不存在
		servLog.Warn("Work Count Not Found In Cache: ", userId)
		// query from db.
		if err := DB.Where("user_id = ?", userId).Count(&workCount).Error; err != nil {
			servLog.Error("QueryWorkCountFromUserId err", err)
			return workCount, err
		}
		servLog.Info("QueryWorkCountFromUserId success")

		// 写入redis缓存
		workCountJson, err := json.Marshal(workCount)
		if err != nil {
			return 0, err
		}
		err = RDB.Set(ctx, strconv.FormatInt(userId, 10), workCountJson, 0).Err()
		if err != nil {
			servLog.Error(err)
			return 0, nil
		}
		return workCount, nil
	}
	// 缓存存在
	servLog.Info("Work Count Get From Cache: ", cacheWorkCount)
	err = json.Unmarshal([]byte(cacheWorkCount), &workCount)
	if err != nil {
		return 0, err
	}
	return workCount, nil
}
