package service

import (
	"context"
	"simple-douyin/kitex_gen/message"
	"simple-douyin/service/message/dal"
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
	result := dal.DB.Create(&newMsg)
	if result.Error != nil || result.RowsAffected == 0 {
		resp.StatusCode = 57007
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "发送失败"
		return nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
