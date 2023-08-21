package service

import (
	"context"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/message"
	"simple-douyin/service/message/dal"

	servLog "github.com/prometheus/common/log"
)

func MessageGetUserLastMessage(ctx context.Context, req *message.MessageGetUserLastMessageRequest, resp *message.MessageGetUserLastMessageResponse) (err error) {
	UserID := req.UserId     // 当前用户ID
	ToUserID := req.ToUserId // 聊天对象ID
	servLog.Info("Message Chat Get: ", req)
	//实际业务
	var messageDb dal.Message

	//去数据库双方最后一条消息 create_at最大的那条

	err = dal.DB.Where("user_id=? and to_user_id=?", UserID, ToUserID).Or("user_id=? and to_user_id=?", ToUserID, UserID).Order("create_at desc").First(&messageDb).Error
	if err != nil {
		return err
	}

	var message common.Message
	message.Id = int64(messageDb.ID)
	message.ToUserId = messageDb.ToUserID
	message.UserId = messageDb.FromUserID
	message.Content = messageDb.Msg
	resp.Message = &message
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
