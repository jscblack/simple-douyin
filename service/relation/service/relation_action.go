package service

import (
	"context"
	"simple-douyin/kitex_gen/relation"
	"simple-douyin/service/relation/dal"

	servLog "github.com/prometheus/common/log"
)

func RelationAdd(ctx context.Context, req *relation.RelationAddRequest, resp *relation.RelationAddResponse) (err error) {
	servLog.Info("Relation Add Get: ", req)
	// 实际业务
	UserID := req.UserId
	ToUserID := req.ToUserId
	result := dal.DB.Create(&dal.Relation{UserID: UserID, ToUserID: ToUserID})
	if result.Error != nil || result.RowsAffected == 0 {
		resp.StatusCode = 57010
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "关注失败"
		servLog.Error("关注失败")
		return nil
	}
	//修改redis缓存
	dal.RDSUpdate(ctx, UserID, 1, 1)
	dal.RDSUpdate(ctx, ToUserID, 1, 2)

	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}

func RelationRemove(ctx context.Context, req *relation.RelationRemoveRequest, resp *relation.RelationRemoveResponse) (err error) {
	servLog.Info("Relation Remove Get: ", req)
	// 实际业务
	UserID := req.UserId
	ToUserID := req.ToUserId
	result := dal.DB.Where("user_id=? and to_user_id=?", UserID, ToUserID).Delete(&dal.Relation{})
	if result.Error != nil || result.RowsAffected == 0 {
		resp.StatusCode = 57010
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "取消关注失败"
		servLog.Error("取消关注失败")
		return nil
	}
	//修改redis缓存
	dal.RDSUpdate(ctx, UserID, -1, 1)
	dal.RDSUpdate(ctx, ToUserID, -1, 2)
	resp.StatusCode = 0
	resp.StatusMsg = nil
	return nil
}
