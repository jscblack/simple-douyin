package main

import (
	"context"
	pong "simple-douyin/kitex_gen/pong"
	"simple-douyin/pkg/constant"
	"time"

	servLog "github.com/sirupsen/logrus"
	"github.com/redis/go-redis/v9"
)

// PongServiceImpl implements the last service interface defined in the IDL.
type PongServiceImpl struct{}

// Pong implements the PongServiceImpl interface.
func (s *PongServiceImpl) Pong(ctx context.Context, req *pong.PingReq) (resp *pong.PongResp, err error) {
	// TODO: Your code here...

	RDB := redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
	})
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	// 执行flushdb命令，清空当前数据库中的所有key
	result := RDB.FlushAll(ctx)
	if result.Err() != nil {
		servLog.Error(result.Err())
		panic(result.Err())
	}
	resp = &pong.PongResp{
		PongTime: req.PingTime + " " + time.Now().String() + " redis flushAll success",
	}
	return resp, nil
}
