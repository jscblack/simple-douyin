package service

import (
	"context"
	"encoding/json"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/user/dal"
	"strconv"

	servLog "github.com/prometheus/common/log"
)

func UserInfo(ctx context.Context, req *user.UserInfoRequest, resp *user.UserInfoResponse) error {
	// 实际业务
	servLog.Info("User Info Get: ", *req.UserId, req.ToUserId)

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
		// 获取计数器信息
		comUser.FollowCount = new(int64)
		comUser.FollowerCount = new(int64)
		comUser.FavoriteCount = new(int64)
		comUser.TotalFavorited = new(int64)
		comUser.WorkCount = new(int64)
		// 获取关注信息
		if req.UserId == nil {
			comUser.IsFollow = false
		} else {
			comUser.IsFollow = false // TODO: 从redis中获取是否关注
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
		return nil
	}
	// 缓存存在
	servLog.Info("User Info Get From Cache: ", cacheUser)
	// 返回数据
	resp.User = &common.User{}
	err = json.Unmarshal([]byte(cacheUser), resp.User)
	if err != nil {
		return err
	}

	// 获取关注信息
	if req.UserId == nil {
		resp.User.IsFollow = false
	} else {
		resp.User.IsFollow = false // TODO: 从redis中获取是否关注
	}

	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
