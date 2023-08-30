package dal

import (
	"context"
	"simple-douyin/pkg/constant"
	"time"

	videoDal "simple-douyin/service/publish/dal"
	userDal "simple-douyin/service/user/dal"

	"github.com/redis/go-redis/v9"
	servLog "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var VDB *redis.Client

// 数据库表结构
type Comment struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"` //视频评论id
	UserID    int64          `gorm:"index" json:"user_id"`               //评论用户id
	VideoID   int64          `gorm:"index" json:"video_id"`              //所评论的视频id
	Content   string         `gorm:"type:varchar(256)" json:"content"`   //评论内容
	CreatedAt time.Time      //AutoCreateTime
	UpdatedAt time.Time      //AutoUpdateTime
	DeletedAt gorm.DeletedAt `gorm:"index"` //AutoDeleteTime
	User      userDal.User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Video     videoDal.Video `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// redis缓存结构
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
	err = DB.AutoMigrate(&Comment{})
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
	// For Redis
	VDB = redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
		DB:       constant.VideoRDB,
	})
	_, err = VDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
}
