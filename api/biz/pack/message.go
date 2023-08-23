package pack

import (
	"context"
	"strconv"

	bizMessage "simple-douyin/api/biz/model/message"
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
	//TODO
	// bizResp.MessageList = messageListPack(ctx, rpcResp.MessageList)
	return nil
}

//

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
