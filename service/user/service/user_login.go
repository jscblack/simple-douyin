package service

import (
	"context"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/user/dal"

	servLog "github.com/prometheus/common/log"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(ctx context.Context, req *user.UserLoginRequest, resp *user.UserLoginResponse) error {
	// 实际业务
	servLog.Info("User Login Get: ", req)
	// 检索数据库查询hash密码
	dalUser := &dal.User{
		Name: req.Username,
	}
	servLog.Info("dalUser: ", dalUser)
	result := dal.DB.Where(dalUser).Take(&dalUser)
	if result.Error != nil || result.RowsAffected == 0 {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "用户不存在"
		servLog.Error("用户不存在")
		return nil
	}
	servLog.Info("dalUser: ", dalUser)
	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(dalUser.Password), []byte(req.Password))
	if err != nil {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "密码错误"
		servLog.Error("密码错误")
		return nil
	}
	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.UserId = dalUser.ID
	return nil
}
