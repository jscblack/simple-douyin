package service

import (
	"context"
	"encoding/json"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/user/dal"
	"strconv"

	servLog "github.com/prometheus/common/log"
)

func UpdateUserCounter(ctx context.Context, req *user.UpdateUserCounterRequest, resp *user.UpdateUserCounterResponse) error {
	servLog.Info("Update User Counter Get: ", req)
	// 该操作期望更新redis中的计数器
	// 但并不一定所有用户均在redis中（比如新用户，冷启动等）
	// 调用一次user_info接口，如果redis中没有该用户，则会将该用户信息加载到redis中
	// 这样就可以保证redis中一定有该用户的计数器
	infoReq := &user.UserInfoRequest{
		ToUserId: req.UserId,
	}
	infoResp := &user.UserInfoResponse{}
	err := UserInfo(ctx, infoReq, infoResp)
	if err != nil {
		return err
	}
	comUser := infoResp.User
	if req.Counter == "FollowCount" {
		if req.Increment {
			*comUser.FollowCount++
		} else {
			*comUser.FollowCount--
		}
	} else if req.Counter == "FollowerCount" {
		if req.Increment {
			*comUser.FollowerCount++
		} else {
			*comUser.FollowerCount--
		}

	} else if req.Counter == "FavoriteCount" {
		if req.Increment {
			*comUser.FavoriteCount++
		} else {
			*comUser.FavoriteCount--
		}
	} else if req.Counter == "TotalFavorited" {
		if req.Increment {
			*comUser.TotalFavorited++
		} else {
			*comUser.TotalFavorited--
		}
	} else if req.Counter == "WorkCount" {
		if req.Increment {
			*comUser.WorkCount++
		} else {
			*comUser.WorkCount--
		}
	} else {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "Counter not found"
		return nil
	}
	// 更新redis
	comUserJson, err := json.Marshal(comUser)
	if err != nil {
		return err
	}
	err = dal.RDB.Set(ctx, strconv.FormatInt(req.UserId, 10), comUserJson, 0).Err()
	if err != nil {
		servLog.Error(err)
	}
	// 返回数据
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
