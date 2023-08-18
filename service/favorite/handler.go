package main

import (
	"context"
	favorite "simple-douyin/kitex_gen/favorite"
	service "simple-douyin/service/favorite/service"

	servLog "github.com/prometheus/common/log"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAddAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAddAction(ctx context.Context, req *favorite.FavoriteAddActionRequest) (resp *favorite.FavoriteAddActionResponse, err error) {
	// TODO: Your code here...
	// 不要忘记更新redis中的数据
	// 取出redis的dal.UserCounter，更新FavorCount和FavoredCount。如果对应的量为-1，请直接略过，如果不为-1，请+1
	// 取出redis的dal.VideoCounter，更新FavoredCount
	// 如果redis中没有对应的counter，直接略过更新redis的过程
	// 先更新mysql，再更新redis，特别需要UserCounter中有两个量，可能值为-1，表示未初始化
	return
}

// FavoriteDelAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteDelAction(ctx context.Context, req *favorite.FavoriteDelActionRequest) (resp *favorite.FavoriteDelActionResponse, err error) {
	// TODO: Your code here...
	// 不要忘记更新redis中的数据
	// 取出redis的dal.UserCounter，更新FavorCount和FavoredCount
	// 取出redis的dal.VideoCounter，更新FavoredCount
	// 如果redis中没有对应的counter，直接略过更新redis的过程
	// 先更新mysql，再更新redis，特别需要UserCounter中有两个量，可能值为-1，表示未初始化
	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// UserFavorCount implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UserFavorCount(ctx context.Context, req *favorite.UserFavorCountRequest) (resp *favorite.UserFavorCountResponse, err error) {
	resp = new(favorite.UserFavorCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.UserFavorCount(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57004
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}

// VideoFavoredCount implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) VideoFavoredCount(ctx context.Context, req *favorite.VideoFavoredCountRequest) (resp *favorite.VideoFavoredCountResponse, err error) {
	resp = new(favorite.VideoFavoredCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.VideoFavoredCount(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57004
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}

// UserFavoredCount implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) UserFavoredCount(ctx context.Context, req *favorite.UserFavoredCountRequest) (resp *favorite.UserFavoredCountResponse, err error) {
	resp = new(favorite.UserFavoredCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.UserFavoredCount(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57004
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}

// IsFavor implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) IsFavor(ctx context.Context, req *favorite.IsFavorRequest) (resp *favorite.IsFavorResponse, err error) {
	resp = new(favorite.IsFavorResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.IsFavored(ctx, req, resp)
	if err != nil {
		resp.StatusCode = 57004
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	// 后处理返回结果
	// ...
	return resp, nil
}
