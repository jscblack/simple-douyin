package main

import (
	"context"
	message "simple-douyin/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, request *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageSend implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageSend(ctx context.Context, request *message.MessageSendRequest) (resp *message.MessageSendResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageGetUserLastMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageGetUserLastMessage(ctx context.Context, request *message.MessageGetUserLastMessageRequest) (resp *message.MessageGetUserLastMessageResponse, err error) {
	// TODO: Your code here...
	return
}
