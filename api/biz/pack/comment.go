package pack

import (
	"context"
	bizComment "simple-douyin/api/biz/model/comment"
	bizCommon "simple-douyin/api/biz/model/common"
	kiteComment "simple-douyin/kitex_gen/comment"
	kiteCommon "simple-douyin/kitex_gen/common"
	"strconv"
)

// CommentAdd .
func CommentAddUnpack(ctx context.Context, bizReq *bizComment.CommentActionRequest, rpcReq *kiteComment.CommentAddActionRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.VideoId = bizReq.VideoID
	rpcReq.CommentText = bizReq.CommentText
	return nil
}

// CommentAddpack .
func CommentAddpack(ctx context.Context, rpcResp *kiteComment.CommentAddActionResponse, bizResp *bizComment.CommentActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	if rpcResp.Comment != nil {
		bizResp.Comment = commentPack(ctx, rpcResp.Comment)
	}
	return nil
}

// CommentDel .
func CommentDelUnpack(ctx context.Context, bizReq *bizComment.CommentActionRequest, rpcReq *kiteComment.CommentDelActionRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.CommentId = *bizReq.CommentID
	return nil
}

// CommentDelpack .
func CommentDelpack(ctx context.Context, rpcResp *kiteComment.CommentDelActionResponse, bizResp *bizComment.CommentActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	return nil
}

// CommentList .
func CommentListUnpack(ctx context.Context, bizReq *bizComment.CommentListRequest, rpcReq *kiteComment.CommentListRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.UserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.VideoId = bizReq.VideoID
	return nil
}

// CommentListpack .
func CommentListpack(ctx context.Context, rpcResp *kiteComment.CommentListResponse, bizResp *bizComment.CommentListResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	if rpcResp.CommentList == nil {
		bizResp.CommentList = make([]*bizCommon.Comment, 0, len(rpcResp.CommentList))
		return nil
	}
	for _, rpcComment := range rpcResp.CommentList {
		bizResp.CommentList = append(bizResp.CommentList, commentPack(ctx, rpcComment))
	}
	return nil
}

func commentPack(ctx context.Context, rpcComment *kiteCommon.Comment) *bizCommon.Comment {
	// rpcComment -> bizComment
	bizComment := new(bizCommon.Comment)
	bizComment.ID = rpcComment.Id
	bizComment.User = UserPack(ctx, rpcComment.User)
	bizComment.Content = rpcComment.Content
	bizComment.CreateDate = rpcComment.CreateDate
	return bizComment
}
