package service

import (
	"context"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/message"
	"simple-douyin/service/message/dal"
	"strconv"
	"time"

	servLog "github.com/sirupsen/logrus"
)

func MessageChat(ctx context.Context, req *message.MessageChatRequest, resp *message.MessageChatResponse) (err error) {
	servLog.Info("Message Chat Get: ", req)
	//实际业务
	var messageListResp []*common.Message
	var messageListDb []*dal.Message
	// req.PreMsgTime is Unixmilli
	preTime := time.UnixMilli(req.PreMsgTime)
	servLog.Info("PreMsgTime: ", preTime)
	//去数据库查询CreateAt在req.PreMsgTime之后的消息 且是两个人之间的 结果全部列存入messageListDb
	err = dal.DB.Where("from_user_id=? and to_user_id=? and created_at>?", req.UserId, req.ToUserId, preTime).Or("from_user_id=? and to_user_id=? and created_at>?", req.ToUserId, req.UserId, preTime).Order("created_at").Find(&messageListDb).Error
	if err != nil {
		return err
	}
	// 后处理messageListDb，将其中时间与preTime过于接近的给他排除
	//将messageListDb转换为messageListResp
	for _, messageDb := range messageListDb {
		// servLog.Info("check", messageDb.CreatedAt, preTime)

		if (messageDb.CreatedAt.UnixMilli() - preTime.UnixMilli()) <= 1500 {
			servLog.Info("same found", messageDb.CreatedAt, preTime)
			continue
		}
		var message common.Message
		message.Id = int64(messageDb.ID)
		message.ToUserId = messageDb.ToUserID
		message.UserId = messageDb.FromUserID
		message.Content = messageDb.Msg
		message.CreateTime = new(string)
		*message.CreateTime = strconv.FormatInt(messageDb.CreatedAt.UnixMilli(), 10)
		messageListResp = append(messageListResp, &message)
	}
	servLog.Info("Message Chat Find: ", len(messageListResp))
	resp.Messages = messageListResp
	return nil
}
