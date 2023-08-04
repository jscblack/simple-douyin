package service

import (
	"context"
	"strconv"

	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/service/favorite/dal"

	servLog "github.com/prometheus/common/log"
)

// 内部rpc
func UserFavorCount(ctx context.Context, req *favorite.UserFavorCountRequest, resp *favorite.UserFavorCountResponse) (err error) {
	// 点赞数
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FavorCount == nil {
		resp.FavorCount = new(int64)
	}
	result := dal.DB.Model(&dal.Favorite{}).Where("user_id=?", req.UserId).Count(resp.FavorCount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func VideoFavorCount(ctx context.Context, req *favorite.VideoFavoredCountRequest, resp *favorite.VideoFavoredCountResponse) (err error) {
	// 视频被点赞数
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FavoredCount == nil {
		resp.FavoredCount = new(int64)
	}
	result := dal.DB.Model(&dal.Favorite{}).Where("video_id=?", req.VideoId).Count(resp.FavoredCount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UserFavoredCount(ctx context.Context, req *favorite.UserFavoredCountRequest, resp *favorite.UserFavoredCountResponse) (err error) {
	// 用户被点赞数
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FavoredCount == nil {
		resp.FavoredCount = new(int64)
	}
	result := dal.DB.Model(&dal.Favorite{}).Where("author_id=?", req.UserId).Count(resp.FavoredCount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func IsFavored(ctx context.Context, req *favorite.IsFavorRequest, resp *favorite.IsFavorResponse) (err error) {
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.IsFavorite = true
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
			servLog.Error("redis set error: ", err)
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
