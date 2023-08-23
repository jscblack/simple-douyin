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
	EtcdAddressWithPort = "172.19.0.1" + ":2379"
	PostgresDSN         = "host=172.19.0.1 user=simple_douyin password=1qaz0plm dbname=simple_douyin_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	RedisAddress        = "172.19.0.1" + ":6379"
	RedisPassword       = "1qaz0plm"
	ServiceAddress      = "127.0.0.1"

	PingServiceName     = "ping_pong"
	PingServicePort     = "7000"
	UserServiceName     = "user"
	UserServicePort     = "7001"
	FeedServiceName     = "feed"
	FeedServicePort     = "7002"
	PublishServiceName  = "publish"
	PublishServicePort  = "7003"
	FavoriteServiceName = "favorite"
	FavoriteServicePort = "7004"
	CommentServiceName  = "comment"
	CommentServicePort  = "7005"
	RelationServiceName = "relation"
	RelationServicePort = "7006"
	MessageServiceName  = "message"
	MessageServicePort  = "7007"

	MaxFeedNum   = 30
	MaxListNum   = 30
	MaxVideoSize = 52428800 // 50MB

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
