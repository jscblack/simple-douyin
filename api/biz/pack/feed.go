package pack

import (
	bizFeed "simple-douyin/api/biz/model/feed"
	kitexFeed "simple-douyin/kitex_gen/feed"
)

// FeedPack kitexResp -> bizResp
func FeedPack(kitexResp *kitexFeed.FeedResponse) (*bizFeed.FeedResponse, error) {
	return &bizFeed.FeedResponse{
		StatusCode: kitexResp.StatusCode,
		StatusMsg:  kitexResp.StatusMsg,
		VideoList:  ToVideos(kitexResp.VideoList),
		NextTime:   kitexResp.NextTime,
	}, nil
}
