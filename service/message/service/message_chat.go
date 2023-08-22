package service

import (
	"context"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/message"
	"simple-douyin/service/message/dal"
	"time"

	servLog "github.com/prometheus/common/log"
)

func MessageChat(ctx context.Context, req *message.MessageChatRequest, resp *message.MessageChatResponse) (err error) {
	servLog.Info("Message Chat Get: ", req)
	//实际业务
	var messageListResp []*common.Message
	var messageListDb []*dal.Message
	//int64转time.Time
	preTime := time.Unix(req.PreMsgTime, 0)
	//去数据库查询CreateAt在req.PreMsgTime之后的消息 且是两个人之间的 结果全部列存入messageListDb
	err = dal.DB.Where("create_at>?", preTime).Where("user_id=? and to_user_id=?", req.UserId, req.ToUserId).Or("user_id=? and to_user_id=?", req.ToUserId, req.UserId).Find(&messageListDb).Error
	if err != nil {
		return err
	}
	//将messageListDb转换为messageListResp
	for _, messageDb := range messageListDb {
		var message common.Message
		message.Id = int64(messageDb.ID)
		message.ToUserId = messageDb.ToUserID
		message.UserId = messageDb.FromUserID
		message.Content = messageDb.Msg
		messageListResp = append(messageListResp, &message)
	}
	resp.Messages = messageListResp
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
