package pack

import (
	bizPublish "simple-douyin/api/biz/model/publish"
	kitexPublish "simple-douyin/kitex_gen/publish"
)

// PublishActionPack kitexResp -> bizResp
func PublishActionPack(kitexResp *kitexPublish.PublishActionResponse) (*bizPublish.PublishActionResponse, error) {
	return &bizPublish.PublishActionResponse{
		StatusCode: kitexResp.StatusCode,
		StatusMsg:  kitexResp.StatusMsg,
	}, nil
}

// PublishListPack kitexResp -> bizResp
func PublishListPack(kitexResp *kitexPublish.PublishListResponse) (*bizPublish.PublishListResponse, error) {
	return &bizPublish.PublishListResponse{
		StatusCode: kitexResp.StatusCode,
		StatusMsg:  kitexResp.StatusMsg,
		VideoList:  ToVideos(kitexResp.VideoList),
	}, nil
}
