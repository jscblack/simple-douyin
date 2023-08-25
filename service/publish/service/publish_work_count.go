package service

import (
	"context"
	servLog "github.com/sirupsen/logrus"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/service/publish/dal"
)

func PublishWorkCount(ctx context.Context, req *publish.PublishWorkCountRequest) (resp *publish.PublishWorkCountResponse, err error) {
	// query video count from db according to userId.
	servLog.Info("Accept request: ", req)

	workCount, err := dal.QueryWorkCountFromUserId(ctx, req.UserId)
	if err != nil {
		servLog.Error("QueryWorkCountFromUserId err", err)
		return nil, err
	}
	resp = publish.NewPublishWorkCountResponse()
	resp.WorkCount = &workCount
	return
}
