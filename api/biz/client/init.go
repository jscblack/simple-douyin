package client

import apiLog "github.com/prometheus/common/log"

func InitClient() {
	InitPingClient()
	InitUserClient()
	InitFeedClient()
	InitPublishClient()
	InitFavoriteClient()
	InitCommentClient()
	InitRelationClient()
	InitMessaggeClient()
	apiLog.Info("All RPC client initialized")
}
