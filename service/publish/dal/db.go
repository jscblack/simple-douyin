package dal

import (
	"context"
	servLog "github.com/sirupsen/logrus"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"simple-douyin/pkg/constant"
	"time"
)

var DB *gorm.DB
var RDB *redis.Client

// User 数据库表结构
type User struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`                  //用户唯一标志符号
	Name            string         `gorm:"type:varchar(128);not null;unique;index" json:"name"` //用户名
	Password        string         `gorm:"type:varchar(128);not null" json:"password"`          //用户密码HMAC
	Avatar          string         `json:"avatar"`                                              //用户头像
	BackgroundImage string         `json:"background_image"`                                    //用户背景图
	Signature       string         `json:"signature"`                                           //用户签名
	CreatedAt       time.Time      //AutoCreateTime
	UpdatedAt       time.Time      //AutoUpdateTime
	DeletedAt       gorm.DeletedAt `gorm:"index"` //AutoDeleteTime
}

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
	User      User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
