package service

import (
	"context"
	"simple-douyin/kitex_gen/comment"
	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/comment/client"
	"simple-douyin/service/comment/dal"
	"sync"

	servLog "github.com/sirupsen/logrus"
)

func CommentList(ctx context.Context, req *comment.CommentListRequest, resp *comment.CommentListResponse) (err error) {
	// 查询评论列表
	// 从pgsql中查询
	var dalComments []dal.Comment
	result := dal.DB.Model(&dal.Comment{}).Where("video_id = ?", req.VideoId).Find(&dalComments)
	if result.Error != nil {
		return result.Error
	}
	// 补全User信息
	resp.CommentList = make([]*common.Comment, len(dalComments))
	var mu sync.Mutex
	var wg sync.WaitGroup
	for idx, dalComment := range dalComments {
		wg.Add(1)
		go func(idx int, dalComment dal.Comment) {
			defer wg.Done()
			userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
				UserId:   &req.UserId,
				ToUserId: dalComment.UserID,
			})
			if err != nil {
				servLog.Error(err)
				return
			}
			mu.Lock()
			resp.CommentList[idx] = &common.Comment{
				Id:      dalComment.ID,
				User:    userResp.User,
				Content: dalComment.Content,
				// yyyy-mm-dd
				CreateDate: dalComment.CreatedAt.Format("2006-01-02"),
			}
			mu.Unlock()
		}(idx, dalComment)

		// userResp, err := client.UserClient.UserInfo(ctx, &user.UserInfoRequest{
		// 	UserId:   &req.UserId,
		// 	ToUserId: dalComment.UserID,
		// })
		// if err != nil {
		// 	servLog.Error(err)
		// 	return err
		// }
		// resp.CommentList = append(resp.CommentList, &common.Comment{
		// 	Id:      dalComment.ID,
		// 	User:    userResp.User,
		// 	Content: dalComment.Content,
		// 	// yyyy-mm-dd
		// 	CreateDate: dalComment.CreatedAt.Format("2006-01-02"),
		// })
	}
	wg.Wait() // Wait for all goroutines to complete
	return nil
}
