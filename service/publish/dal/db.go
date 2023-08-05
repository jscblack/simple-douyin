package dal

import (
	"context"
	servLog "github.com/prometheus/common/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple-douyin/pkg/constant"
)

var DB *gorm.DB
var _ *redis.Client

// Video 数据库表结构
type Video struct {
	gorm.Model
	UserId     int64  `json:"user_id"`
	PlayUrl    string `json:"play_url"`
	CoverUrl   string `json:"cover_url"`
	CreateTime int64  `gorm:"default:0"`
	Title      string `json:"title"`
}

// Init 初始化，创建数据库连接
func Init(ctx context.Context) {
	var err error
	DB, err = gorm.Open(postgres.Open(constant.PostgresDSN),
		&gorm.Config{},
	)
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	err = DB.AutoMigrate(&Video{})
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	// For Redis
	//RDB = redis.NewClient(&redis.Options{
	//	Addr:     constant.RedisAddress,
	//	Password: constant.RedisPassword, // 没有密码，默认值
	//	DB:       1,                      // DB 0 for User ; DB 1 for Video
	//})
	//_, err = RDB.Ping(ctx).Result()
	//if err != nil {
	//	servLog.Error(err)
	//	panic(err)
	//}
}
