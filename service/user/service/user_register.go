package service

import (
	"context"
	"simple-douyin/kitex_gen/user"
	"simple-douyin/service/user/dal"

	servLog "github.com/prometheus/common/log"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(ctx context.Context, req *user.UserRegisterRequest, resp *user.UserRegisterResponse) error {
	// 实际业务
	servLog.Info("User Register Get: ", req)
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	dalUser := &dal.User{
		Name:     req.Username,
		Password: string(bcryptPassword),
	}

	result := dal.DB.Create(&dalUser)
	if result.Error != nil || result.RowsAffected == 0 {
		resp.StatusCode = 57001
		if resp.StatusMsg == nil {
			resp.StatusMsg = new(string)
		}
		*resp.StatusMsg = "用户名已存在"
		servLog.Error("用户名已存在")
		return nil
	}

	resp.StatusCode = 0
	resp.StatusMsg = nil
	resp.UserId = dalUser.ID
	return nil
}
