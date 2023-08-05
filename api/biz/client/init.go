package client

import apiLog "github.com/prometheus/common/log"

func InitClient() {
	InitPingClient()
	InitUserClient()
	InitFeedClient()
	InitPublishClient()
	apiLog.Info("All RPC client initialized")
}
