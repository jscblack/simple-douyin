package pack

import (
	"context"
	bizCommon "simple-douyin/api/biz/model/common"
	bizUser "simple-douyin/api/biz/model/user"
	kiteCommon "simple-douyin/kitex_gen/common"
	kiteUser "simple-douyin/kitex_gen/user"
	"strconv"
)

// UserRegisterUnpack .
func UserRegisterUnpack(ctx context.Context, bizReq *bizUser.UserRegisterRequest, rpcReq *kiteUser.UserRegisterRequest) error {
	// bizReq -> rpcReq
	rpcReq.Username = bizReq.Username
	rpcReq.Password = bizReq.Password
	return nil
}

// UserRegisterPack .
func UserRegisterPack(ctx context.Context, rpcResp *kiteUser.UserRegisterResponse, bizResp *bizUser.UserRegisterResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.UserID = rpcResp.UserId
	return nil
}

// UserLoginUnpack .
func UserLoginUnpack(ctx context.Context, bizReq *bizUser.UserLoginRequest, rpcReq *kiteUser.UserLoginRequest) error {
	// bizReq -> rpcReq
	rpcReq.Username = bizReq.Username
	rpcReq.Password = bizReq.Password
	return nil
}

// UserLoginPack .
func UserLoginPack(ctx context.Context, rpcResp *kiteUser.UserLoginResponse, bizResp *bizUser.UserLoginResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.UserID = rpcResp.UserId
	return nil
}

// UserInfoUnpack .
func UserInfoUnpack(ctx context.Context, bizReq *bizUser.UserInfoRequest, rpcReq *kiteUser.UserInfoRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId = new(int64)
		*rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.ToUserId = bizReq.UserID

	return nil
}

// UserInfoPack .
func UserInfoPack(ctx context.Context, rpcResp *kiteUser.UserInfoResponse, bizResp *bizUser.UserInfoResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.User = userPack(ctx, rpcResp.User)
	return nil
}

func userPack(ctx context.Context, rpcUser *kiteCommon.User) *bizCommon.User {
	// rpcUser -> bizUser
	bizUser := new(bizCommon.User)
	bizUser.ID = rpcUser.Id
	bizUser.Name = rpcUser.Name
	bizUser.FollowCount = rpcUser.FollowCount
	bizUser.FollowerCount = rpcUser.FollowerCount
	bizUser.IsFollow = rpcUser.IsFollow
	bizUser.Avatar = rpcUser.Avatar
	bizUser.BackgroundImage = rpcUser.BackgroundImage
	bizUser.Signature = rpcUser.Signature
	bizUser.TotalFavorited = rpcUser.TotalFavorited
	bizUser.WorkCount = rpcUser.WorkCount
	bizUser.FavoriteCount = rpcUser.FavoriteCount
	return bizUser
}
