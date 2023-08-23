package pack

import (
	"context"
	bizPublish "simple-douyin/api/biz/model/publish"
	kitexPublish "simple-douyin/kitex_gen/publish"
)

// PublishActionPack kitexResp -> bizResp
func PublishActionPack(ctx context.Context, kitexResp *kitexPublish.PublishActionResponse) (*bizPublish.PublishActionResponse, error) {
	return &bizPublish.PublishActionResponse{
		StatusCode: kitexResp.StatusCode,
		StatusMsg:  kitexResp.StatusMsg,
	}, nil
}

// PublishListPack kitexResp -> bizResp
func PublishListPack(ctx context.Context, kitexResp *kitexPublish.PublishListResponse) (*bizPublish.PublishListResponse, error) {
	return &bizPublish.PublishListResponse{
		StatusCode: kitexResp.StatusCode,
		StatusMsg:  kitexResp.StatusMsg,
		VideoList:  ToVideos(ctx, kitexResp.VideoList),
	}, nil
}
