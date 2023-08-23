package service

import (
	"context"
	"encoding/json"
	"simple-douyin/kitex_gen/comment"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/comment/client"
	"simple-douyin/service/comment/dal"
	"strconv"

	servLog "github.com/prometheus/common/log"
)

func CommentAdd(ctx context.Context, req *comment.CommentAddActionRequest, resp *comment.CommentAddActionResponse) (err error) {
	dalComment := dal.Comment{
		UserID:  req.UserId,
		VideoID: req.VideoId,
		Content: *req.CommentText,
	}
	result := dal.DB.Model(&dal.Comment{}).Create(&dalComment)
	if result.Error != nil {
		return result.Error
	}
	// 更新视频的被评论总数CommentCount
	cacheVideoCounter, err := dal.VDB.Get(ctx, strconv.FormatInt(dalComment.VideoID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalVideoCounter := dal.VideoCounter{}
		err = json.Unmarshal([]byte(cacheVideoCounter), &dalVideoCounter)
		if err != nil {
			return err
		}
		if dalVideoCounter.CommentCount != -1 {
			dalVideoCounter.CommentCount++
		}
		videoCounterJson, err := json.Marshal(dalVideoCounter)
		if err != nil {
			return err
		}
		err = dal.VDB.Set(ctx, strconv.FormatInt(dalComment.VideoID, 10), videoCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
		UserId:   &req.UserId,
		ToUserId: dalComment.UserID,
	})
	if err != nil {
		servLog.Error(err)
		return err
	}
	resp.Comment = &common.Comment{
		Id:      dalComment.ID,
		Content: dalComment.Content,
		User:    userResp.User,
		// yyyy-mm-dd
		CreateDate: dalComment.CreatedAt.Format("2006-01-02"),
	}
	return nil
}

func CommentDel(ctx context.Context, req *comment.CommentDelActionRequest, resp *comment.CommentDelActionResponse) (err error) {
	dalComment := dal.Comment{
		ID: req.CommentId,
	}
	// 校验是否为评论者
	result := dal.DB.Model(&dal.Comment{}).Where(&dalComment).First(&dalComment)
	if result.Error != nil {
		return result.Error
	}
	if dalComment.UserID != req.UserId {
		resp.StatusCode = 57005
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "not the comment owner"
		return nil
	}
	// 删除评论
	result = dal.DB.Model(&dal.Comment{}).Delete(&dalComment)
	if result.Error != nil {
		return result.Error
	}
	// 更新视频的CommentCount
	cacheVideoCounter, err := dal.VDB.Get(ctx, strconv.FormatInt(dalComment.VideoID, 10)).Result()
	if err == nil {
		//缓存存在，更新
		dalVideoCounter := dal.VideoCounter{}
		err = json.Unmarshal([]byte(cacheVideoCounter), &dalVideoCounter)
		if err != nil {
			return err
		}
		if dalVideoCounter.CommentCount != -1 {
			dalVideoCounter.CommentCount--
		}
		videoCounterJson, err := json.Marshal(dalVideoCounter)
		if err != nil {
			return err
		}
		err = dal.VDB.Set(ctx, strconv.FormatInt(dalComment.VideoID, 10), videoCounterJson, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
