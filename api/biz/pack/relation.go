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
	rpcReq.UserId = bizReq.UserID
	return nil
}

func RelationFriendListPack(ctx context.Context, bizReq *bizRelation.RelationFriendListResponse, rpcReq *kiteRelation.RelationFriendListResponse) error {
	// rpcResp -> bizResp
	bizReq.StatusMsg = rpcReq.StatusMsg
	bizReq.StatusCode = rpcReq.StatusCode
	//TODO
	return nil
}
