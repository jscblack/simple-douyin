package dal

import (
	"context"
	"simple-douyin/pkg/constant"
	"time"

	servLog "github.com/prometheus/common/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var RDB *redis.Client

// 数据库表结构
type Comment struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"` //视频评论id
	UserID    int64          `gorm:"index" json:"user_id"`               //评论用户id
	VideoID   int64          `gorm:"index" json:"video_id"`              //所评论的视频id
	Content   string         `gorm:"type:varchar(256)" json:"content"`   //评论内容
	CreatedAt time.Time      //AutoCreateTime
	UpdatedAt time.Time      //AutoUpdateTime
	DeletedAt gorm.DeletedAt `gorm:"index"` //AutoDeleteTime
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
	err = DB.AutoMigrate(&Comment{})
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	// // For Redis
	// RDB = redis.NewClient(&redis.Options{
	// 	Addr:     constant.RedisAddress,
	// 	Password: constant.RedisPassword, // 没有密码，默认值
	// 	DB:       *,                      // DB * for *
	// })
	// _, err = RDB.Ping(ctx).Result()
	// if err != nil {
	// 	servLog.Error(err)
	// 	panic(err)
	// }
}
