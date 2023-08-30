package service

import (
	"context"
	"sync"

	"simple-douyin/kitex_gen/common"
	"simple-douyin/kitex_gen/favorite"
	"simple-douyin/kitex_gen/publish"
	"simple-douyin/service/favorite/client"
	"simple-douyin/service/favorite/dal"

	servLog "github.com/sirupsen/logrus"
)

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest, resp *favorite.FavoriteListResponse) (err error) {
	// 查询点赞列表
	// 从pgsql中查询
	var dalFavs []dal.Favorite
	result := dal.DB.Model(&dal.Favorite{}).Where("user_id = ?", req.UserId).Find(&dalFavs)
	if result.Error != nil {
		return result.Error
	}
	// 补全Video信息

	// for _, dalFav := range dalFavs {
	// 	publishResp, err := client.PublishClient.PublishVideoInfo(ctx, &publish.PublishVideoInfoRequest{
	// 		UserId:  &req.FromUserId,
	// 		VideoId: dalFav.VideoID,
	// 	})
	// 	if err != nil {
	// 		servLog.Error(err)
	// 		return err
	// 	}
	// 	resp.VideoList = append(resp.VideoList, publishResp.Video)
	// }
	resp.VideoList = make([]*common.Video, len(dalFavs))
	var mu sync.Mutex
	var wg sync.WaitGroup
	for idx, dalFav := range dalFavs {
		wg.Add(1)
		go func(idx int, dalFav dal.Favorite) {
			defer wg.Done()
			publishResp, err := client.PublishClient.PublishVideoInfo(ctx, &publish.PublishVideoInfoRequest{
				UserId:  &req.FromUserId,
				VideoId: dalFav.VideoID,
			})
			if err != nil {
				servLog.Error(err)
				return
			}
			mu.Lock()
			resp.VideoList[idx] = publishResp.Video
			mu.Unlock()
		}(idx, dalFav)
	}
	wg.Wait() // Wait for all goroutines to complete

	return nil
}
