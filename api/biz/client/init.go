package client

import apiLog "github.com/prometheus/common/log"

func InitClient() {
	InitPingClient()
	InitUserClient()
	InitPublishClient()
	apiLog.Info("All RPC client initialized")
}
