package dal

import (
	"context"
	"simple-douyin/pkg/constant"

	userDal "simple-douyin/service/user/dal"

	servLog "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 数据库表结构
// 同时缓存于Redis中
// key: user_id to_user_id 中间用空格隔开
type Message struct {
	gorm.Model
	// gorm.Model equals:
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	FromUserID int64        `gorm:"index" json:"user_id"`
	ToUserID   int64        `gorm:"index" json:"to_user_id"`
	Msg        string       `json:"msg"`
	FromUser   userDal.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ToUser     userDal.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// 初始化，创建数据库连接
func Init(ctx context.Context) {
	var err error
	DB, err = gorm.Open(postgres.Open(constant.PostgresDSN),
		&gorm.Config{},
	)
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	err = DB.AutoMigrate(&Message{})
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
}
