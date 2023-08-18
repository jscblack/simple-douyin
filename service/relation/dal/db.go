package dal

import (
	"context"
	"simple-douyin/pkg/constant"

	servLog "github.com/prometheus/common/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

// 数据库表结构
// 同时缓存于Redis中
// key: user_id to_user_id 中间用空格隔开
type Relation struct {
	gorm.Model
	// gorm.Model equals:
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID   int64 `gorm:"index" json:"user_id"`
	ToUserID int64 `gorm:"index" json:"to_user_id"`
}

// Redis缓存结构
// key: user_id
type RelationCounter struct {
	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
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
	err = DB.AutoMigrate(&Relation{})
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	// For Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
		DB:       constant.RelationRDB,
	})
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
}
