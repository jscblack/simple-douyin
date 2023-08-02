package service

import (
	"context"
	"simple-douyin/kitex_gen/relation"
)

// 下面的三个函数是内部rpc，需要接入redis
func RelationFollowCount(ctx context.Context, req *relation.RelationFollowCountRequest, resp *relation.RelationFollowCountRequest) error {
	return nil
}

func RelationFollowerCount(ctx context.Context, req *relation.RelationFollowerCountRequest, resp *relation.RelationFollowerCountRequest) error {
	return nil
}

func RelationIsFollow(ctx context.Context, req *relation.RelationIsFollowRequest, resp *relation.RelationIsFollowRequest) error {
	return nil
}
