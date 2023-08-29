package constant

const (
	// NoteTableName           = "note"
	// UserTableName           = "user"
	// SecretKey               = "secret key"
	// IdentityKey             = "id"
	// Total                   = "total"
	// Notes                   = "notes"
	// NoteID                  = "note_id"
	// ApiServiceName          = "demoapi"
	// NoteServiceName         = "demonote"
	// UserServiceName         = "demouser"
	// CPURateLimit    float64 = 80.0
	// DefaultLimit            = 10
	EtcdAddressWithPort = "simple-douyin-etcd" + ":2379"
	PostgresDSN         = "host=simple-douyin-postgres user=simple_douyin password=1qaz0plm dbname=simple_douyin_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	RedisAddress        = "simple-douyin-redis" + ":6379"
	RedisPassword       = "1qaz0plm"
	ServiceAddress      = "127.0.0.1"

	PingServiceName = "ping_pong"
	PingServicePort = "7000"

	UserServiceName      = "user"
	UserServicePort      = "7001"
	UserServerTracerPort = ":9102"
	UserServerTracerPath = "/userserver"
	UserClientTracerPort = ":9103"
	UserClientTracerPath = "/userclient"

	FeedServiceName      = "feed"
	FeedServicePort      = "7002"
	FeedServerTracerPort = ":9104"
	FeedServerTracerPath = "/feedserver"
	FeedClientTracerPort = ":9105"
	FeedClientTracerPath = "/feedclient"

	PublishServiceName      = "publish"
	PublishServicePort      = "7003"
	PublishServerTracerPort = ":9106"
	PublishServerTracerPath = "/publishserver"
	PublishClientTracerPort = ":9107"
	PublishClientTracerPath = "/publishclient"

	FavoriteServiceName      = "favorite"
	FavoriteServicePort      = "7004"
	FavoriteServerTracerPort = ":9108"
	FavoriteServerTracerPath = "/favoriteserver"
	FavoriteClientTracerPort = ":9109"
	FavoriteClientTracerPath = "/favoriteclient"

	CommentServiceName      = "comment"
	CommentServicePort      = "7005"
	CommentServerTracerPort = ":9110"
	CommentServerTracerPath = "/commentserver"
	CommentClientTracerPort = ":9111"
	CommentClientTracerPath = "/commentclient"

	RelationServiceName      = "relation"
	RelationServicePort      = "7006"
	RelationServerTracerPort = ":9112"
	RelationServerTracerPath = "/relationserver"
	RelationClientTracerPort = ":9113"
	RelationClientTracerPath = "/relationclient"

	MessageServiceName      = "message"
	MessageServicePort      = "7007"
	MessageServerTracerPort = ":9114"
	MessageServerTracerPath = "/messageserver"
	MessageClientTracerPort = ":9115"
	MessageClientTracerPath = "/messageclient"

	MaxFeedNum   = 30
	MaxListNum   = 30
	MaxVideoSize = 100 * 1024 * 1024 // 100MB

	UserRDB     = 0
	PublishRDB  = 1
	FavoriteRDB = 2
	VideoRDB    = 3
	RelationRDB = 4

	// for upy oss
	Bucket      = "simple-douyin-oos"
	Operator    = "simpledouyin"
	UpyPassword = "xbTwBs6LDtbqVSIXmim6QYu2xGSdM6Jr"
)
