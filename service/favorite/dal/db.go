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
type Favorite struct {
	gorm.Model
	// gorm.Model equals:
	// ID        uint `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID   int64 `gorm:"index" json:"user_id"`
	VideoID  int64 `gorm:"index" json:"video_id"`
	AuthorID int64 `gorm:"index" json:"author_id"` // 作者ID，以加快查找被赞数
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
		DB:       3,                      // DB 3 for Favorite
	})
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		servLog.Error(err)
		panic(err)
	}
}
