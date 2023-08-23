package service

import (
	"context"
	"encoding/json"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/kitex_gen/relation"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/user/client"
	"simple-douyin/service/user/dal"
	"strconv"

	servLog "github.com/prometheus/common/log"
)

func fillUserInfo(ctx context.Context, comUser *common.User, req *user.UserInfoRequest) error {
	// 通过rpc补全用户信息
	// 获取计数器信息
	followCountResp, err := client.RelationClient.RelationFollowCount(ctx, &relation.RelationFollowCountRequest{
		UserId: comUser.Id,
	})
	if err != nil {
		return err
	}
	followerCountResp, err := client.RelationClient.RelationFollowerCount(ctx, &relation.RelationFollowerCountRequest{
		UserId: comUser.Id,
	})
	if err != nil {
		return err
	}
	favoriteCountResp, err := client.FavoriteClient.UserFavorCount(ctx, &favorite.UserFavorCountRequest{
		UserId: comUser.Id,
	})
	if err != nil {
		return err
	}
	totalFavoritedResp, err := client.FavoriteClient.UserFavoredCount(ctx, &favorite.UserFavoredCountRequest{
		UserId: comUser.Id,
	})
	if err != nil {
		return err
	}
	workCountResp, err := client.PublishClient.PublishWorkCount(ctx, &publish.PublishWorkCountRequest{
		UserId: comUser.Id,
	})
	if err != nil {
		return err
	}
	comUser.FollowCount = followCountResp.FollowCount
	comUser.FollowerCount = followerCountResp.FollowerCount
	comUser.FavoriteCount = favoriteCountResp.FavorCount
	comUser.TotalFavorited = totalFavoritedResp.FavoredCount
	comUser.WorkCount = workCountResp.WorkCount
	// comUser.FollowCount = new(int64)
	// *comUser.FollowCount = *followCountResp.FollowCount
	// comUser.FollowerCount = new(int64)
	// *comUser.FollowerCount = *followerCountResp.FollowerCount
	// comUser.FavoriteCount = new(int64)
	// *comUser.FavoriteCount = *favoriteCountResp.FavorCount
	// comUser.TotalFavorited = new(int64)
	// *comUser.TotalFavorited = *totalFavoritedResp.FavoredCount
	// comUser.WorkCount = new(int64)
	// *comUser.WorkCount = *workCountResp.WorkCount

	// 获取关注信息
	if req.UserId == nil {
		comUser.IsFollow = false
	} else {
		isFollowResp, err := client.RelationClient.RelationIsFollow(ctx, &relation.RelationIsFollowRequest{
			UserId:   *req.UserId,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			return err
		}
		comUser.IsFollow = isFollowResp.IsFollow
	}
	return nil
}

func UserInfo(ctx context.Context, req *user.UserInfoRequest, resp *user.UserInfoResponse) error {
	// 实际业务
	servLog.Info("User Info Get: ", req.UserId, req.ToUserId)

	// 检查redis缓存
	cacheUser, err := dal.RDB.Get(ctx, strconv.FormatInt(req.ToUserId, 10)).Result()
	if err != nil {
		// 缓存不存在
		servLog.Warn("User Info Not Fountd In Cahce: ", req.ToUserId)
		dalUser := &dal.User{
			ID: req.ToUserId,
		}
		result := dal.DB.Where(dalUser).Take(&dalUser)
		if result.Error != nil || result.RowsAffected == 0 {
			resp.StatusCode = 57001
			if resp.StatusMsg == nil {
				resp.StatusMsg = new(string)
			}
			*resp.StatusMsg = "用户不存在"
			resp.User = new(common.User)
			servLog.Error("用户不存在")
			return nil
		}
		comUser := &common.User{
			Id:              dalUser.ID,
			Name:            dalUser.Name,
			Avatar:          &dalUser.Avatar,
			BackgroundImage: &dalUser.BackgroundImage,
			Signature:       &dalUser.Signature,
		}
		// 写入redis缓存
		comUserJson, err := json.Marshal(comUser)
		if err != nil {
			return err
		}
		err = dal.RDB.Set(ctx, strconv.FormatInt(req.ToUserId, 10), comUserJson, 0).Err()
		if err != nil {
			servLog.Error(err)
		}
		// 返回数据
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.User = comUser
		err = fillUserInfo(ctx, resp.User, req)
		if err != nil {
			servLog.Error(err)
		}
		return nil
	}
	// 缓存存在
	servLog.Info("User Info Get From Cache: ", cacheUser)
	// 返回数据
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.User = &common.User{}
	err = json.Unmarshal([]byte(cacheUser), resp.User)
	if err != nil {
		return err
	}
	err = fillUserInfo(ctx, resp.User, req)
	if err != nil {
		servLog.Error(err)
	}
	return nil
}
