package pack

import (
	"context"
	bizCommon "simple-douyin/api/biz/model/common"
	bizFav "simple-douyin/api/biz/model/favorite"
	kiteCommon "simple-douyin/kitex_gen/common"
	kiteFav "simple-douyin/kitex_gen/favorite"
	"strconv"
)

// FavoriteAdd .
func FavoriteAddUnpack(ctx context.Context, bizReq *bizFav.FavoriteActionRequest, rpcReq *kiteFav.FavoriteAddActionRequest) error {
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

// FavoriteAddpack .
func FavoriteAddpack(ctx context.Context, rpcResp *kiteFav.FavoriteAddActionResponse, bizResp *bizFav.FavoriteActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	return nil
}

// FavoriteDel .
func FavoriteDelUnpack(ctx context.Context, bizReq *bizFav.FavoriteActionRequest, rpcReq *kiteFav.FavoriteDelActionRequest) error {
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

// FavoriteDelpack .
func FavoriteDelpack(ctx context.Context, rpcResp *kiteFav.FavoriteDelActionResponse, bizResp *bizFav.FavoriteActionResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	return nil
}

// FavoriteList .
func FavoriteListUnpack(ctx context.Context, bizReq *bizFav.FavoriteListRequest, rpcReq *kiteFav.FavoriteListRequest) error {
	// bizReq -> rpcReq
	if bizReq.Token != "" {
		var err error
		rpcReq.FromUserId, err = strconv.ParseInt(bizReq.Token, 10, 64)
		if err != nil {
			return err
		}
	}
	rpcReq.UserId = bizReq.UserID
	return nil
}

// FavoriteListpack .
func FavoriteListpack(ctx context.Context, rpcResp *kiteFav.FavoriteListResponse, bizResp *bizFav.FavoriteListResponse) error {
	// rpcResp -> bizResp
	bizResp.StatusCode = rpcResp.StatusCode
	bizResp.StatusMsg = rpcResp.StatusMsg
	for _, rpcVideo := range rpcResp.VideoList {
		bizVideo := videoPack(ctx, rpcVideo)
		bizResp.VideoList = append(bizResp.VideoList, bizVideo)
	}
	return nil
}

func videoPack(ctx context.Context, rpcVideo *kiteCommon.Video) *bizCommon.Video {
	// rpcVideo-> bizVideo
	bizVideo := new(bizCommon.Video)
	bizVideo.ID = rpcVideo.Id
	bizVideo.Author = UserPack(ctx, rpcVideo.Author)
	bizVideo.PlayURL = rpcVideo.PlayUrl
	bizVideo.CoverURL = rpcVideo.CoverUrl
	bizVideo.FavoriteCount = rpcVideo.FavoriteCount
	bizVideo.CommentCount = rpcVideo.CommentCount
	bizVideo.IsFavorite = rpcVideo.IsFavorite
	bizVideo.Title = rpcVideo.Title
	return bizVideo
}
