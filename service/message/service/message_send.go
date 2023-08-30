package service

import (
	"context"
	"simple-douyin/kitex_gen/message"
	"simple-douyin/kitex_gen/relation"
	"simple-douyin/service/message/client"
	"simple-douyin/service/message/dal"

	servLog "github.com/sirupsen/logrus"
)

func MessageSend(ctx context.Context, req *message.MessageSendRequest, resp *message.MessageSendResponse) (err error) {
	//业务逻辑
	FromUserID := req.UserId
	ToUserID := req.ToUserId
	Content := req.Content

	newMsg := &dal.Message{
		FromUserID: FromUserID,
		ToUserID:   ToUserID,
		Msg:        Content,
	}
	// 需要确保FromUserID和ToUserID确实是好友关系
	relResp, err := client.RelationClient.RelationIsFollow(ctx, &relation.RelationIsFollowRequest{
		UserId:   FromUserID,
		ToUserId: ToUserID,
	})
	if err != nil {
		servLog.Error("RelationIsFollow error: ", err)
		return err
	}
	if !relResp.IsFollow {
		resp.StatusCode = 57007
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "发送失败"
		return nil
	} else {
		relResp, err = client.RelationClient.RelationIsFollow(ctx, &relation.RelationIsFollowRequest{
			UserId:   ToUserID,
			ToUserId: FromUserID,
		})
		if err != nil {
			servLog.Error("RelationIsFollow error: ", err)
			return err
		}
		if !relResp.IsFollow {
			resp.StatusCode = 57007
			if resp.StatusMsg == nil {
				resp.StatusMsg = new(string)
			}
			*resp.StatusMsg = "发送失败"
			return nil
		}
	}
	servLog.Info("message send: ", newMsg)
	result := dal.DB.Create(&newMsg)
	if result.Error != nil || result.RowsAffected == 0 {
		resp.StatusCode = 57007
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "发送失败"
		return nil
	}
	return nil
}
