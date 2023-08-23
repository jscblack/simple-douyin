package main

import (
	"context"
	relation "simple-douyin/kitex_gen/relation"
	"simple-douyin/service/relation/service"

	servLog "github.com/prometheus/common/log"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAdd implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAdd(ctx context.Context, request *relation.RelationAddRequest) (resp *relation.RelationAddResponse, err error) {
	// TODO: Your code here...
	// 不要忘记更新redis中的FollowCount和FollowerCount
	// 取出redis的dal.RelationCounter，更新FollowCount和FollowerCount
	// 如果redis中没有对应的counter，处理逻辑参见./service/relation_count.go
	// 先更新mysql，再更新redis，特别需要RelationCounter中有两个量，可能值为-1，表示未初始化

	resp = new(relation.RelationAddResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationAdd(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}

	return resp, nil
}

// RelationRemove implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationRemove(ctx context.Context, request *relation.RelationRemoveRequest) (resp *relation.RelationRemoveResponse, err error) {
	// TODO: Your code here...
	// 不要忘记更新redis中的FollowCount和FollowerCount
	// 取出redis的dal.RelationCounter，更新FollowCount和FollowerCount
	// 如果redis中没有对应的counter，处理逻辑参见./service/relation_count.go
	// 先更新mysql，再更新redis，特别需要RelationCounter中有两个量，可能值为-1，表示未初始化
	resp = new(relation.RelationRemoveResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationRemove(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}

	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, request *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.RelationFollowListResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationFollowList(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, request *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.RelationFollowerListResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationFollowerList(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, request *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.RelationFriendListResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationFriendList(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = err.Error()
		servLog.Error(err)
		return resp, nil
	}
	return resp, nil
}

// RelationFollowCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowCount(ctx context.Context, request *relation.RelationFollowCountRequest) (resp *relation.RelationFollowCountResponse, err error) {
	resp = new(relation.RelationFollowCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationFollowCount(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
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

// RelationFollowerCount implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerCount(ctx context.Context, request *relation.RelationFollowerCountRequest) (resp *relation.RelationFollowerCountResponse, err error) {
	resp = new(relation.RelationFollowerCountResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationFollowerCount(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
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

// RelationIsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationIsFollow(ctx context.Context, request *relation.RelationIsFollowRequest) (resp *relation.RelationIsFollowResponse, err error) {
	resp = new(relation.RelationIsFollowResponse)
	// 前处理校验请求
	// ...
	// 实际业务
	err = service.RelationIsFollow(ctx, request, resp)
	if err != nil {
		resp.StatusCode = 57006
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
