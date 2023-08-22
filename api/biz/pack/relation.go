package pack

import (
	"context"
	bizRelation "simple-douyin/api/biz/model/relation"
	kiteRelation "simple-douyin/kitex_gen/relation"
	"strconv"
)

func RelationAddUnpack(ctx context.Context, bizReq *bizRelation.RelationActionRequest, rpcReq *kiteRelation.RelationAddRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.ToUserId = bizReq.ToUserID
	return nil
}
func RelationAddPack(ctx context.Context, rpcResp *kiteRelation.RelationAddResponse, bizResp *bizRelation.RelationActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.StatusCode = rpcResp.StatusCode
	return nil
}

func RelationRemoveUnpack(ctx context.Context, bizReq *bizRelation.RelationActionRequest, rpcReq *kiteRelation.RelationRemoveRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.ToUserId = bizReq.ToUserID
	return nil
}

func RelationRemovePack(ctx context.Context, rpcResp *kiteRelation.RelationRemoveResponse, bizResp *bizRelation.RelationActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.StatusCode = rpcResp.StatusCode
	return nil
}

// RelationFollowListUnpack .
func RelationFollowListUnpack(ctx context.Context, bizReq *bizRelation.RelationFollowListRequest, rpcReq *kiteRelation.RelationFollowListRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.FromUserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.UserId = bizReq.UserID
	return nil
}
func RelationFollowListPack(ctx context.Context, rpcResp *kiteRelation.RelationFollowListResponse, bizResp *bizRelation.RelationFollowListResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.UserList = userListPack(ctx, rpcResp.FollowList)
	return nil
}

// RelationFollowerListUnpack .
func RelationFollowerListUnpack(ctx context.Context, bizReq *bizRelation.RelationFollowerListRequest, rpcReq *kiteRelation.RelationFollowerListRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.FromUserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.UserId = bizReq.UserID
	return nil
}

func RelationFollowerListPack(ctx context.Context, rpcResp *kiteRelation.RelationFollowerListResponse, bizResp *bizRelation.RelationFollowerListResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.UserList = userListPack(ctx, rpcResp.FollowerList)
	return nil
}

func RelationFriendListUnpack(ctx context.Context, bizReq *bizRelation.RelationFriendListRequest, rpcReq *kiteRelation.RelationFriendListRequest) error {
	// bizReq -> rpcReq

	if bizReq.Token != "" {
		var err error
		rpcReq.FromUserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}

	rpcReq.UserId = bizReq.UserID
	return nil
}

// type FriendUser struct{
// 	Id              int64   `thrift:"id,1,required" frugal:"1,required,i64" json:"id"`
// 	Name            string  `thrift:"name,2,required" frugal:"2,required,string" json:"name"`
// 	FollowCount     *int64  `thrift:"follow_count,3,optional" frugal:"3,optional,i64" json:"follow_count,omitempty"`
// 	FollowerCount   *int64  `thrift:"follower_count,4,optional" frugal:"4,optional,i64" json:"follower_count,omitempty"`
// 	IsFollow        bool    `thrift:"is_follow,5,required" frugal:"5,required,bool" json:"is_follow"`
// 	Avatar          *string `thrift:"avatar,6,optional" frugal:"6,optional,string" json:"avatar,omitempty"`
// 	BackgroundImage *string `thrift:"background_image,7,optional" frugal:"7,optional,string" json:"background_image,omitempty"`
// 	Signature       *string `thrift:"signature,8,optional" frugal:"8,optional,string" json:"signature,omitempty"`
// 	TotalFavorited  *int64  `thrift:"total_favorited,9,optional" frugal:"9,optional,i64" json:"total_favorited,omitempty"`
// 	WorkCount       *int64  `thrift:"work_count,10,optional" frugal:"10,optional,i64" json:"work_count,omitempty"`
// 	FavoriteCount   *int64  `thrift:"favorite_count,11,optional" frugal:"11,optional,i64" json:"favorite_count,omitempty"`
// 	Message         *string `thrift:"message,12,optional" frugal:"12,optional,string" json:"message,omitempty"`
// 	MsgType         int64   `thrift:"msg_type,13,required" frugal:"13,required,i64" json:"msg_type"`
// }

func RelationFriendListPack(ctx context.Context, rpcReq *kiteRelation.RelationFriendListResponse, bizReq *bizRelation.RelationFriendListResponse) error {
	// rpcResp -> bizResp
	bizReq.StatusMsg = rpcReq.StatusMsg
	bizReq.StatusCode = rpcReq.StatusCode
	//TODO
	var friendList []*bizRelation.FriendUser
	for _, friend := range rpcReq.FriendList {
		friendList = append(friendList, &bizRelation.FriendUser{
			ID:              friend.Id,
			Name:            friend.Name,
			FollowCount:     friend.FollowCount,
			FollowerCount:   friend.FollowerCount,
			IsFollow:        friend.IsFollow,
			Avatar:          friend.Avatar,
			BackgroundImage: friend.BackgroundImage,
			Signature:       friend.Signature,
			TotalFavorited:  friend.TotalFavorited,
			WorkCount:       friend.WorkCount,
			FavoriteCount:   friend.FavoriteCount,
			Message:         friend.Message,
			MsgType:         friend.MsgType,
		})
	}
	bizReq.UserList = friendList
	return nil
}
