package service

import (
	"context"
	"encoding/json"
	"strconv"

	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/service/favorite/client"
	"simple-douyin/service/favorite/dal"

	servLog "github.com/sirupsen/logrus"
)

func FavoriteAddAction(ctx context.Context, req *favorite.FavoriteAddActionRequest, resp *favorite.FavoriteAddActionResponse) (err error) {
	// 点赞操作
	publishResp, err := client.PublishClient.PublishVideoInfo(ctx, &publish.PublishVideoInfoRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		servLog.Error(err)
		return err
	}
	dalFav := dal.Favorite{
		UserID:   req.UserId,
		VideoID:  req.VideoId,
		AuthorID: publishResp.Video.Author.Id,
	}
	// 如果存在则不创建
	result := dal.DB.Model(&dal.Favorite{}).Where(&dalFav).First(&dalFav)
	if result.Error == nil {
		// 存在该记录，不创建
		resp.StatusCode = 57004
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "already favored"
		return nil
	}
	// 写入pgsql
	result = dal.DB.Model(&dal.Favorite{}).Create(&dalFav)
	if result.Error != nil {
		return result.Error
	}
	// 更新点赞者计数器
	cacheUserCounter, err := dal.RDB.Get(ctx, strconv.FormatInt(dalFav.UserID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalUserCounter := dal.UserCounter{}
		err = json.Unmarshal([]byte(cacheUserCounter), &dalUserCounter)
		if err != nil {
			return err
		}
		if dalUserCounter.FavorCount != -1 {
			dalUserCounter.FavorCount++
		}
		userCounterJson, err := json.Marshal(dalUserCounter)
		if err != nil {
			return err
		}
		err = dal.RDB.Set(ctx, strconv.FormatInt(dalFav.UserID, 10), userCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	// 更新被点赞者计数器
	cacheUserCounter, err = dal.RDB.Get(ctx, strconv.FormatInt(dalFav.AuthorID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalUserCounter := dal.UserCounter{}
		err = json.Unmarshal([]byte(cacheUserCounter), &dalUserCounter)
		if err != nil {
			return err
		}
		if dalUserCounter.FavoredCount != -1 {
			dalUserCounter.FavoredCount++
		}
		userCounterJson, err := json.Marshal(dalUserCounter)
		if err != nil {
			return err
		}
		err = dal.RDB.Set(ctx, strconv.FormatInt(dalFav.AuthorID, 10), userCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	// 更新视频计数器
	cacheVideoCounter, err := dal.VDB.Get(ctx, strconv.FormatInt(dalFav.VideoID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalVideoCounter := dal.VideoCounter{}
		err = json.Unmarshal([]byte(cacheVideoCounter), &dalVideoCounter)
		if err != nil {
			return err
		}
		if dalVideoCounter.FavoredCount != -1 {
			dalVideoCounter.FavoredCount++
		}
		videoCounterJson, err := json.Marshal(dalVideoCounter)
		if err != nil {
			return err
		}
		err = dal.VDB.Set(ctx, strconv.FormatInt(dalFav.VideoID, 10), videoCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	// 更新被点赞关系
	keyStr := strconv.FormatInt(dalFav.UserID, 10) + " " + strconv.FormatInt(dalFav.VideoID, 10)
	err = dal.RDB.Set(ctx, keyStr, "1", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func FavoriteDelAction(ctx context.Context, req *favorite.FavoriteDelActionRequest, resp *favorite.FavoriteDelActionResponse) (err error) {
	// 取消点赞操作
	dalFav := dal.Favorite{
		UserID:  req.UserId,
		VideoID: req.VideoId,
	}
	// 先查询
	result := dal.DB.Model(&dal.Favorite{}).Where(&dalFav).First(&dalFav)
	if result.Error != nil {
		// 不存在，抛出异常
		return result.Error
	}
	// 从pgsql中删除
	result = dal.DB.Model(&dal.Favorite{}).Where(&dalFav).Delete(&dalFav)
	if result.Error != nil {
		return result.Error
	}

	// 更新点赞者计数器
	cacheUserCounter, err := dal.RDB.Get(ctx, strconv.FormatInt(dalFav.UserID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalUserCounter := dal.UserCounter{}
		err = json.Unmarshal([]byte(cacheUserCounter), &dalUserCounter)
		if err != nil {
			return err
		}
		if dalUserCounter.FavorCount != -1 {
			dalUserCounter.FavorCount--
		}
		userCounterJson, err := json.Marshal(dalUserCounter)
		if err != nil {
			return err
		}
		err = dal.RDB.Set(ctx, strconv.FormatInt(dalFav.UserID, 10), userCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	// 更新被点赞者计数器
	cacheUserCounter, err = dal.RDB.Get(ctx, strconv.FormatInt(dalFav.AuthorID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalUserCounter := dal.UserCounter{}
		err = json.Unmarshal([]byte(cacheUserCounter), &dalUserCounter)
		if err != nil {
			return err
		}
		if dalUserCounter.FavoredCount != -1 {
			dalUserCounter.FavoredCount--
		}
		userCounterJson, err := json.Marshal(dalUserCounter)
		if err != nil {
			return err
		}
		err = dal.RDB.Set(ctx, strconv.FormatInt(dalFav.AuthorID, 10), userCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	// 更新视频计数器
	cacheVideoCounter, err := dal.VDB.Get(ctx, strconv.FormatInt(dalFav.VideoID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalVideoCounter := dal.VideoCounter{}
		err = json.Unmarshal([]byte(cacheVideoCounter), &dalVideoCounter)
		if err != nil {
			return err
		}
		if dalVideoCounter.FavoredCount != -1 {
			dalVideoCounter.FavoredCount--
		}
		videoCounterJson, err := json.Marshal(dalVideoCounter)
		if err != nil {
			return err
		}
		err = dal.VDB.Set(ctx, strconv.FormatInt(dalFav.VideoID, 10), videoCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	// 更新被点赞关系
	keyStr := strconv.FormatInt(dalFav.UserID, 10) + " " + strconv.FormatInt(dalFav.VideoID, 10)
	err = dal.RDB.Set(ctx, keyStr, "0", 0).Err()
	if err != nil {
		return err
	}
	return nil
}
