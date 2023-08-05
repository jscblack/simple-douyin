package service

import (
	"context"

	"simple-douyin/kitex_gen/comment"
	"simple-douyin/service/comment/dal"
)

// 这是内部rpc
func CommentCount(ctx context.Context, req *comment.CommentCountRequest, resp *comment.CommentCountResponse) (err error) {
	// 该rpc不需要接入redis，直接从mysql中查询，因为其会被其他rpc缓存
	resp.StatusCode = 0
	resp.StatusMsg = nil
	if resp.CommentCount == nil {
		resp.CommentCount = new(int64)
	}
	result := dal.DB.Model(&dal.Comment{}).Where("video_id=?", req.VideoId).Count(resp.CommentCount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
