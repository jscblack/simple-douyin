package service

import (
	"context"
	"encoding/json"
	"strconv"

	"simple-douyin/kitex_gen/comment"
	"simple-douyin/service/comment/dal"

	servLog "github.com/sirupsen/logrus"
)

// 这是内部rpc
func CommentCount(ctx context.Context, req *comment.CommentCountRequest, resp *comment.CommentCountResponse) (err error) {
	// 视频被点赞数，use VDB
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.CommentCount == nil {
		resp.CommentCount = new(int64)
	}
	*resp.CommentCount = 0                                      // 初始化
	keyStr := strconv.FormatInt(req.VideoId, 10)                // int64转string
	cacheVideoCounter, err := dal.VDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中
		result := dal.DB.Model(&dal.Comment{}).Where("video_id=?", req.VideoId).Count(resp.CommentCount)
		if result.Error != nil {
			return result.Error
		}
		videoCounterJson, err := json.Marshal(&dal.VideoCounter{
			FavoredCount: -1,
			CommentCount: *resp.CommentCount,
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
	if videoCounter.CommentCount == -1 {
		// 对应的count未初始化
		result := dal.DB.Model(&dal.Comment{}).Where("video_id=?", req.VideoId).Count(resp.CommentCount)
		if result.Error != nil {
			return result.Error
		}
		videoCounter.CommentCount = *resp.CommentCount
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
	*resp.CommentCount = videoCounter.CommentCount
	return nil
}
