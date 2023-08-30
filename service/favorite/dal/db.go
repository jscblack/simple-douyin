package dal

import (
	"context"
	"simple-douyin/pkg/constant"

	videoDal "simple-douyin/service/publish/dal"
	userDal "simple-douyin/service/user/dal"

	"github.com/redis/go-redis/v9"
	servLog "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
var VDB *redis.Client

// 数据库表结构
// 同时缓存于Redis中
// key: user_id video_id 中间用空格隔开
type Favorite struct {
	gorm.Model
	// gorm.Model equals:
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID   int64          `gorm:"index" json:"user_id"`
	VideoID  int64          `gorm:"index" json:"video_id"`
	AuthorID int64          `gorm:"index" json:"author_id"` // 作者ID，以加快查找被赞数
	User     userDal.User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Video    videoDal.Video `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Author   userDal.User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Redis缓存结构
// key: user_id
type UserCounter struct {
	FavorCount   int64 `json:"favor_count"`   // 点赞数
	FavoredCount int64 `json:"favored_count"` // 被点赞数
}

// key: video_id
// VideoCounter is used to cache the number of favorites and comments of a video
type VideoCounter struct {
	FavoredCount int64 `json:"favored_count"` // 点赞数
	CommentCount int64 `json:"comment_count"` // 评论数
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
	err = DB.AutoMigrate(&Favorite{})
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	// For Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
		DB:       constant.FavoriteRDB,   // 存放Favorite和UserCounter
	})
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	VDB = redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
		DB:       constant.VideoRDB,      // 存放VideoCounter
	})
	_, err = VDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
}
