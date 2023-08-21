package service

import (
	"context"
	"encoding/json"
	"strconv"

	"simple-douyin/kitex_gen/relation"
	"simple-douyin/service/relation/dal"

	servLog "github.com/prometheus/common/log"
)

// 下面的三个函数是内部rpc
func RelationFollowCount(ctx context.Context, req *relation.RelationFollowCountRequest, resp *relation.RelationFollowCountResponse) error {
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FollowCount == nil {
		resp.FollowCount = new(int64)
	}
	*resp.FollowCount = 0                                          // 初始化
	keyStr := strconv.FormatInt(req.UserId, 10)                    // int64转string
	cacheRealtionCounter, err := dal.RDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中
		result := dal.DB.Model(&dal.Relation{}).Where("user_id=?", req.UserId).Count(resp.FollowCount)
		if result.Error != nil {
			return result.Error
		}
		relationCounterJson, err := json.Marshal(&dal.RelationCounter{
			FollowCount:   *resp.FollowCount,
			FollowerCount: -1,
		})
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, relationCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中, 解析json
	var relationCounter dal.RelationCounter
	err = json.Unmarshal([]byte(cacheRealtionCounter), &relationCounter)
	if err != nil {
		return err
	}
	if relationCounter.FollowCount == -1 {
		// 对应的count未初始化
		result := dal.DB.Model(&dal.Relation{}).Where("user_id=?", req.UserId).Count(resp.FollowCount)
		if result.Error != nil {
			return result.Error
		}
		relationCounter.FollowCount = *resp.FollowCount
		relationCounterJson, err := json.Marshal(relationCounter)
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, relationCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	*resp.FollowCount = relationCounter.FollowCount
	
	return nil
}

func RelationFollowerCount(ctx context.Context, req *relation.RelationFollowerCountRequest, resp *relation.RelationFollowerCountResponse) error {
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.FollowerCount == nil {
		resp.FollowerCount = new(int64)
	}
	*resp.FollowerCount = 0                                        // 初始化
	keyStr := strconv.FormatInt(req.UserId, 10)                    // int64转string
	cacheRealtionCounter, err := dal.RDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中
		result := dal.DB.Model(&dal.Relation{}).Where("to_user_id=?", req.UserId).Count(resp.FollowerCount)
		if result.Error != nil {
			return result.Error
		}
		relationCounterJson, err := json.Marshal(&dal.RelationCounter{
			FollowCount:   -1,
			FollowerCount: *resp.FollowerCount,
		})
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, relationCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中, 解析json
	var relationCounter dal.RelationCounter
	err = json.Unmarshal([]byte(cacheRealtionCounter), &relationCounter)
	if err != nil {
		return err
	}
	if relationCounter.FollowerCount == -1 {
		// 对应的count未初始化
		result := dal.DB.Model(&dal.Relation{}).Where("to_user_id=?", req.UserId).Count(resp.FollowerCount)
		if result.Error != nil {
			return result.Error
		}
		relationCounter.FollowerCount = *resp.FollowerCount
		relationCounterJson, err := json.Marshal(relationCounter)
		if err != nil {
			return err
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, relationCounterJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	*resp.FollowerCount = relationCounter.FollowerCount
	return nil
}

func RelationIsFollow(ctx context.Context, req *relation.RelationIsFollowRequest, resp *relation.RelationIsFollowResponse) error {
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.IsFollow = true
	// 首先从redis中查询
	// int64转string
	keyStr := strconv.FormatInt(req.UserId, 10) + " " + strconv.FormatInt(req.ToUserId, 10)
	cacheRel, err := dal.RDB.Get(ctx, keyStr).Result()
	if err != nil {
		// 不在缓存中
		dalRelation := &dal.Relation{
			UserID:   req.UserId,
			ToUserID: req.ToUserId,
		}
		result := dal.DB.Where(dalRelation).Take(&dalRelation)
		relStr := "1"
		if result.Error != nil || result.RowsAffected == 0 {
			// 不存在关注关系
			relStr = "0"
			resp.IsFollow = false
		}
		// 写入redis缓存
		err = dal.RDB.Set(ctx, keyStr, relStr, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 在缓存中
	if cacheRel == "0" {
		resp.IsFollow = false
	} else {
		resp.IsFollow = true
	}
	return nil
}
