package service

import (
	"context"

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
	resp.VideoList = make([]*common.Video, 0, len(dalFavs))
	for _, dalFav := range dalFavs {
		publishResp, err := client.PublishClient.PublishVideoInfo(ctx, &publish.PublishVideoInfoRequest{
			UserId:  &req.FromUserId,
			VideoId: dalFav.VideoID,
		})
		if err != nil {
			servLog.Error(err)
			return err
		}
		resp.VideoList = append(resp.VideoList, publishResp.Video)
	}
	return nil
}
