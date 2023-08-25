package main

import (
	"context"
	message "simple-douyin/kitex_gen/message"
	"simple-douyin/service/message/service"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	resp = new(message.MessageChatResponse)
	err = service.MessageChat(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57007
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return resp, nil
	}
	return resp, nil
}

// MessageSend implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageSend(ctx context.Context, req *message.MessageSendRequest) (resp *message.MessageSendResponse, err error) {
	resp = new(message.MessageSendResponse)
	err = service.MessageSend(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57007
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return resp, nil
	}
	return resp, nil
}

// MessageGetUserLastMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageGetUserLastMessage(ctx context.Context, req *message.MessageGetUserLastMessageRequest) (resp *message.MessageGetUserLastMessageResponse, err error) {
	resp = new(message.MessageGetUserLastMessageResponse)
	err = service.MessageGetUserLastMessage(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57007
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		return resp, nil
	}
	return resp, nil
}
