package pack

import (
	"context"
	bizFav "simple-douyin/api/biz/model/favorite"
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
