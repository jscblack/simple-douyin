package client

import apiLog "github.com/prometheus/common/log"

func InitClient() {
	InitPingClient()
	InitUserClient()
	InitFeedClient()
	apiLog.Info("All RPC client initialized")
}
