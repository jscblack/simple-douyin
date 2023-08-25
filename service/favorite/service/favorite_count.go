package service

import (
	"context"
	"encoding/json"
	"strconv"

	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/service/favorite/dal"

	servLog "github.com/sirupsen/logrus"
)

// 内部rpc
func UserFavorCount(ctx context.Context, req *favorite.UserFavorCountRequest, resp *favorite.UserFavorCountResponse) (err error) {
	// 点赞数
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FavorCount == nil {
		resp.FavorCount = new(int64)
	}
	*resp.FavorCount = 0                                       // 初始化
	keyStr := strconv.FormatInt(req.UserId, 10)                // int64转string
	cacheUserCounter, err := dal.RDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中
		result := dal.DB.Model(&dal.Favorite{}).Where("user_id=?", req.UserId).Count(resp.FavorCount)
		if result.Error != nil {
			return result.Error
		}
		userCounterJson, err := json.Marshal(&dal.UserCounter{
			FavorCount:   *resp.FavorCount,
			FavoredCount: -1,
		})
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, userCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中, 解析json
	var userCounter dal.UserCounter
	err = json.Unmarshal([]byte(cacheUserCounter), &userCounter)
	if err != nil {
		return err
	}
	if userCounter.FavorCount == -1 {
		// 对应的count未初始化
		result := dal.DB.Model(&dal.Favorite{}).Where("user_id=?", req.UserId).Count(resp.FavorCount)
		if result.Error != nil {
			return result.Error
		}
		userCounter.FavorCount = *resp.FavorCount
		userCounterJson, err := json.Marshal(userCounter)
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, userCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	*resp.FavorCount = userCounter.FavorCount
	return nil
}

func UserFavoredCount(ctx context.Context, req *favorite.UserFavoredCountRequest, resp *favorite.UserFavoredCountResponse) (err error) {
	// 用户被点赞数
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FavoredCount == nil {
		resp.FavoredCount = new(int64)
	}
	*resp.FavoredCount = 0                                     // 初始化
	keyStr := strconv.FormatInt(req.UserId, 10)                // int64转string
	cacheUserCounter, err := dal.RDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中
		result := dal.DB.Model(&dal.Favorite{}).Where("author_id=?", req.UserId).Count(resp.FavoredCount)
		if result.Error != nil {
			return result.Error
		}
		userCounterJson, err := json.Marshal(&dal.UserCounter{
			FavorCount:   -1,
			FavoredCount: *resp.FavoredCount,
		})
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, userCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中, 解析json
	var userCounter dal.UserCounter
	err = json.Unmarshal([]byte(cacheUserCounter), &userCounter)
	if err != nil {
		return err
	}
	if userCounter.FavoredCount == -1 {
		// 对应的count未初始化
		result := dal.DB.Model(&dal.Favorite{}).Where("author_id=?", req.UserId).Count(resp.FavoredCount)
		if result.Error != nil {
			return result.Error
		}
		userCounter.FavoredCount = *resp.FavoredCount
		userCounterJson, err := json.Marshal(userCounter)
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, userCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	*resp.FavoredCount = userCounter.FavoredCount
	return nil
}

func VideoFavoredCount(ctx context.Context, req *favorite.VideoFavoredCountRequest, resp *favorite.VideoFavoredCountResponse) (err error) {
	// 视频被点赞数，use VDB
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FavoredCount == nil {
		resp.FavoredCount = new(int64)
	}
	*resp.FavoredCount = 0                                      // 初始化
	keyStr := strconv.FormatInt(req.VideoId, 10)                // int64转string
	cacheVideoCounter, err := dal.VDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中
		result := dal.DB.Model(&dal.Favorite{}).Where("video_id=?", req.VideoId).Count(resp.FavoredCount)
		if result.Error != nil {
			return result.Error
		}
		videoCounterJson, err := json.Marshal(&dal.VideoCounter{
			FavoredCount: *resp.FavoredCount,
			CommentCount: -1,
		})
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.VDB.Set(ctx, keyStr, videoCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中, 解析json
	var videoCounter dal.VideoCounter
	err = json.Unmarshal([]byte(cacheVideoCounter), &videoCounter)
	if err != nil {
		return err
	}
	if videoCounter.FavoredCount == -1 {
		// 对应的count未初始化
		result := dal.DB.Model(&dal.Favorite{}).Where("video_id=?", req.VideoId).Count(resp.FavoredCount)
		if result.Error != nil {
			return result.Error
		}
		videoCounter.FavoredCount = *resp.FavoredCount
		videoCounterJson, err := json.Marshal(videoCounter)
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.VDB.Set(ctx, keyStr, videoCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	*resp.FavoredCount = videoCounter.FavoredCount
	return nil
}

func IsFavored(ctx context.Context, req *favorite.IsFavorRequest, resp *favorite.IsFavorResponse) (err error) {
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.IsFavorite = true
	if req.UserId == 0 || req.VideoId == 0 {
		resp.IsFavorite = false
		return nil
	}
	// 首先从redis中查询
	// int64转string
	keyStr := strconv.FormatInt(req.UserId, 10) + " " + strconv.FormatInt(req.VideoId, 10)
	cacheRel, err := dal.RDB.Get(ctx, keyStr).Result()
	if err != nil {
		// 不在缓存中
		dalFavorite := &dal.Favorite{
			UserID:  req.UserId,
			VideoID: req.VideoId,
		}
		result := dal.DB.Where(dalFavorite).Take(&dalFavorite)
		relStr := "1"
		if result.Error != nil || result.RowsAffected == 0 {
			// 不存在点赞关系
			relStr = "0"
			resp.IsFavorite = false
		}
		// 存入redis
		err = dal.RDB.Set(ctx, keyStr, relStr, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中
	if cacheRel == "0" {
		resp.IsFavorite = false
	} else {
		resp.IsFavorite = true
	}
	return nil
}
