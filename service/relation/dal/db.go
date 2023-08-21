package dal

import (
	"context"
	"encoding/json"
	"simple-douyin/pkg/constant"
	"strconv"

	apiLog "github.com/prometheus/common/log"
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
		apiLog.Error(err)
		panic(err)
	}
	err = DB.AutoMigrate(&Relation{})
	if err != nil {
		apiLog.Error(err)
		panic(err)
	}
	// For Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     constant.RedisAddress,
		Password: constant.RedisPassword, // 没有密码，默认值
		// DB:       constant.RelationRDB,
		DB: 0,
	})
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		apiLog.Error(err)
		panic(err)
	}
}

func RDSUpdate(ctx context.Context, UserId int64, add int64, Type int32) {
	keyStr := strconv.FormatInt(UserId, 10)                    // int64转string
	cacheRealtionCounter, err := RDB.Get(ctx, keyStr).Result() // 从redis中查询
	if err != nil {
		// 不在缓存中 DO NOTHING
		apiLog.Info("RDSUpdate : not in cache")
		return
	}
	var relationCounter RelationCounter
	err = json.Unmarshal([]byte(cacheRealtionCounter), &relationCounter) // 解析json
	if err != nil {
		apiLog.Error("RDSUpdate : json unmarshal error")
		return
	}
	if Type == 1 { // follow_count
		if relationCounter.FollowCount == -1 {
			apiLog.Info("RDSUpdate : follow_count not init")
			return
		} else {
			relationCounter.FollowCount += add
			relationCounterJson, err := json.Marshal(relationCounter)
			if err != nil {
				apiLog.Error("RDSUpdate : json marshal error")
				return
			}
			// 写入redis缓存
			err = RDB.Set(ctx, keyStr, relationCounterJson, 0).Err()
			if err != nil {
				apiLog.Error("RDSUpdate : redis set error")
				return
			}
		}
	} else { // follower_count
		if relationCounter.FollowerCount == -1 {
			apiLog.Info("RDSUpdate : follower_count not init")
			return
		} else {
			relationCounter.FollowerCount += add
			relationCounterJson, err := json.Marshal(relationCounter)
			if err != nil {
				apiLog.Error("RDSUpdate : json marshal error")
				return
			}
			// 写入redis缓存
			err = RDB.Set(ctx, keyStr, relationCounterJson, 0).Err()
			if err != nil {
				apiLog.Error("RDSUpdate : redis set error")
				return
			}
		}
	}
}
