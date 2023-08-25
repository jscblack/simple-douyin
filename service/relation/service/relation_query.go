package service

import (
	"context"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/message"
	"simple-douyin/kitex_gen/relation"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/relation/client"
	"simple-douyin/service/relation/dal"

	// common
	servLog "github.com/prometheus/common/log"
)

func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest, resp *relation.RelationFollowListResponse) (err error) {
	UserID := req.UserId
	servLog.Info("Relation FollowList Get: ", req)
	// 实际业务
	var followListId []int64
	err = dal.DB.Model(&dal.Relation{}).Where("user_id=?", UserID).Pluck("to_user_id", &followListId).Error
	if err != nil {
		return err
	}
	//根据followListId查询用户信息
	var followListRpc []*common.User
	for _, id := range followListId {
		// get user info
		userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId:   &req.FromUserId,
			ToUserId: id,
		})
		if err != nil {
			return err
		}
		followListRpc = append(followListRpc, userResp.User)
	}
	resp.FollowList = followListRpc
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}

func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest, resp *relation.RelationFollowerListResponse) (err error) {
	UserID := req.UserId
	servLog.Info("Relation FollowedList Get: ", req)
	// 实际业务
	var followerListId []int64
	err = dal.DB.Model(&dal.Relation{}).Where("to_user_id=?", UserID).Pluck("user_id", &followerListId).Error
	if err != nil {
		return err
	}
	//根据followerListId查询用户信息
	var followerListRpc []*common.User
	for _, id := range followerListId {
		// get user info
		userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId:   &req.FromUserId,
			ToUserId: id,
		})
		if err != nil {
			return err
		}
		followerListRpc = append(followerListRpc, userResp.User)
	}
	resp.FollowerList = followerListRpc
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}

func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest, resp *relation.RelationFriendListResponse) (err error) {
	//TODO: 未实现
	UserID := req.UserId
	servLog.Info("Relation FriendList Get: ", req)
	// 实际业务
	var followListId []int64
	var friendListId []int64 //交集
	err = dal.DB.Model(&dal.Relation{}).Where("user_id=?", UserID).Pluck("to_user_id", &followListId).Error
	if err != nil {
		return err
	}

	for _, follow_id := range followListId {
		//如果有 user_id = follow_id and to_user_id = req.UserId
		var relation dal.Relation
		err = dal.DB.Where("user_id=? and to_user_id=?", follow_id, UserID).Limit(1).Find(&relation).Error
		if err != nil {
			return err
		}
		if relation.ID != 0 {
			friendListId = append(friendListId, follow_id)
		}
	}

	//根据friendListId查询用户信息
	var friendListRpc []*relation.FriendUser

	for _, id := range friendListId {
		// get user info
		userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
			UserId:   &UserID,
			ToUserId: id,
		})
		if err != nil {
			return err
		}
		// Id              int64   `thrift:"id,1,required" frugal:"1,required,i64" json:"id"`
		// Name            string  `thrift:"name,2,required" frugal:"2,required,string" json:"name"`
		// FollowCount     *int64  `thrift:"follow_count,3,optional" frugal:"3,optional,i64" json:"follow_count,omitempty"`
		// FollowerCount   *int64  `thrift:"follower_count,4,optional" frugal:"4,optional,i64" json:"follower_count,omitempty"`
		// IsFollow        bool    `thrift:"is_follow,5,required" frugal:"5,required,bool" json:"is_follow"`
		// Avatar          *string `thrift:"avatar,6,optional" frugal:"6,optional,string" json:"avatar,omitempty"`
		// BackgroundImage *string `thrift:"background_image,7,optional" frugal:"7,optional,string" json:"background_image,omitempty"`
		// Signature       *string `thrift:"signature,8,optional" frugal:"8,optional,string" json:"signature,omitempty"`
		// TotalFavorited  *int64  `thrift:"total_favorited,9,optional" frugal:"9,optional,i64" json:"total_favorited,omitempty"`
		// WorkCount       *int64  `thrift:"work_count,10,optional" frugal:"10,optional,i64" json:"work_count,omitempty"`
		// FavoriteCount   *int64  `thrift:"favorite_count,11,optional" frugal:"11,optional,i64" json:"favorite_count,omitempty"`
		// Message         *string `thrift:"message,12,optional" frugal:"12,optional,string" json:"message,omitempty"`
		// MsgType         int64   `thrift:"msg_type,13,required" frugal:"13,required,i64" json:"msg_type"`
		// var friendUser relation.FriendUser
		friendUser := new(relation.FriendUser)
		//new 指针

		friendUser.Id = userResp.User.Id
		friendUser.Name = userResp.User.Name
		friendUser.FollowCount = userResp.User.FollowCount
		friendUser.FollowerCount = userResp.User.FollowerCount
		friendUser.IsFollow = userResp.User.IsFollow
		friendUser.Avatar = userResp.User.Avatar
		friendUser.BackgroundImage = userResp.User.BackgroundImage
		friendUser.Signature = userResp.User.Signature
		friendUser.TotalFavorited = userResp.User.TotalFavorited
		friendUser.WorkCount = userResp.User.WorkCount
		friendUser.FavoriteCount = userResp.User.FavoriteCount

		//调用message服务 GetLastUserMessage

		messageResp, err := client.MessageClient.MessageGetUserLastMessage(ctx, &message.MessageGetUserLastMessageRequest{
			UserId:   UserID,
			ToUserId: id,
		})

		if err != nil {
			return err
		}
		friendUser.Message = &messageResp.Message.Content

		//判断消息类型:
		if messageResp.Message.UserId == UserID {
			friendUser.MsgType = 1
		} else {
			friendUser.MsgType = 2
		}

		friendListRpc = append(friendListRpc, friendUser)

	}
	resp.FriendList = friendListRpc
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
