package client

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	apiLog "github.com/prometheus/common/log"
	bizCommon "simple-douyin/api/biz/model/common"
	bizFeed "simple-douyin/api/biz/model/feed"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/feed"
	"simple-douyin/kitex_gen/feed/feedservice"
	"simple-douyin/pkg/constant"
	"time"
)

var feedClient feedservice.Client // interface from RPC IDL

func InitFeedClient() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddressWithPort})
	if err != nil {
		apiLog.Fatal(err)
	}
	c, err := feedservice.NewClient(
		constant.FeedServiceName,
		// client.WithMiddleware(middleware.CommonMiddleware),
		// client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    // mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		// client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		// client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
	)
	if err != nil {
		apiLog.Fatal(err)
	}
	feedClient = c
	apiLog.Info("Feed client initialized")
}

func Feed(ctx context.Context, req *feed.FeedRequest) (*bizFeed.FeedResponse, error) {
	resp, err := feedClient.Feed(ctx, req)
	if err != nil {
		apiLog.Fatal(err)
	}
	return &bizFeed.FeedResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		VideoList:  ToVideos(resp.VideoList),
		NextTime:   resp.NextTime,
	}, nil
}

func ToVideo(video *common.Video) *bizCommon.Video {
	if video == nil {
		return nil
	}

	author := video.Author
	user := &bizCommon.User{
		ID:            author.Id,
		Name:          author.Name,
		FollowCount:   author.FollowerCount,
		FollowerCount: author.FollowerCount,
		IsFollow:      author.IsFollow,
	}

	return &bizCommon.Video{
		ID:            video.Id,
		Author:        user,
		PlayURL:       video.PlayUrl,
		CoverURL:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
}

func ToVideos(videos []*common.Video) []*bizCommon.Video {
	videoList := make([]*bizCommon.Video, 0)
	for _, video := range videos {
		if v := ToVideo(video); v != nil {
			videoList = append(videoList, v)
		} else {
			apiLog.Fatal("video is nil!")
		}
	}
	return videoList
}
