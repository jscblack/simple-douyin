package dal

import (
	"context"
	"simple-douyin/pkg/constant"
	userDal "simple-douyin/service/user/dal"
	"time"

	"github.com/redis/go-redis/v9"
	servLog "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

// Video 数据库表结构
type Video struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId    int64          `json:"user_id" gorm:"index;"`
	PlayUrl   string         `json:"play_url"`
	CoverUrl  string         `json:"cover_url"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `gorm:"index"` //AutoCreateTime
	UpdatedAt time.Time      //AutoUpdateTime
	DeletedAt gorm.DeletedAt `gorm:"index"` //AutoDeleteTime
	User      userDal.User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	servLog.Info("DB initialized")
	// For Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
		DB:       constant.PublishRDB,    // DB 0 for User ; DB 1 for Video
	})
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	servLog.Info("Redis initialized")
}
