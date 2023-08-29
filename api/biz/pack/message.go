package pack

import (
	"context"
	"strconv"

	bizCommon "simple-douyin/api/biz/model/common"
	bizMessage "simple-douyin/api/biz/model/message"
	kiteCommon "simple-douyin/kitex_gen/common"
	kiteMessage "simple-douyin/kitex_gen/message"
)

func MessageChatUnpack(ctx context.Context, bizReq *bizMessage.MessageChatRequest, rpcReq *kiteMessage.MessageChatRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.ToUserId = bizReq.ToUserID
	rpcReq.PreMsgTime = bizReq.PreMsgTime
	return nil
}

func MessageChatPack(ctx context.Context, rpcResp *kiteMessage.MessageChatResponse, bizResp *bizMessage.MessageChatResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.MessageList = messageListPack(ctx, rpcResp.Messages)
	return nil
}

func MessageSendUnpack(ctx context.Context, bizReq *bizMessage.MessageActionRequest, rpcReq *kiteMessage.MessageSendRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.ToUserId = bizReq.ToUserID
	rpcReq.Content = bizReq.Content
	return nil
}

func MessageSendPack(ctx context.Context, rpcResp *kiteMessage.MessageSendResponse, bizResp *bizMessage.MessageActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusMsg = rpcResp.StatusMsg
	bizResp.StatusCode = rpcResp.StatusCode
	return nil
}

func messageListPack(ctx context.Context, rpcResp []*kiteCommon.Message) []*bizCommon.Message {
	var bizResp []*bizCommon.Message
	for _, rpcMsg := range rpcResp {
		var bizMsg bizCommon.Message
		bizMsg.ID = rpcMsg.Id
		bizMsg.Content = rpcMsg.Content
		bizMsg.FromUserID = rpcMsg.UserId
		bizMsg.ToUserID = rpcMsg.ToUserId
		bizMsg.CreateTime = rpcMsg.CreateTime
		bizResp = append(bizResp, &bizMsg)
	}
	return bizResp
}
