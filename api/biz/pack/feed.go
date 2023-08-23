package pack

import (
	"context"
	bizFeed "simple-douyin/api/biz/model/feed"
	kitexFeed "simple-douyin/kitex_gen/feed"
)

// FeedPack kitexResp -> bizResp
func FeedPack(ctx context.Context, kitexResp *kitexFeed.FeedResponse) (*bizFeed.FeedResponse, error) {
	return &bizFeed.FeedResponse{
		StatusCode: kitexResp.StatusCode,
		StatusMsg:  kitexResp.StatusMsg,
		VideoList:  ToVideos(ctx, kitexResp.VideoList),
		NextTime:   kitexResp.NextTime,
	}, nil
}
