package main

import (
	"context"
	"net"
	message "simple-douyin/kitex_gen/message/messageservice"
	"simple-douyin/pkg/constant"
	"simple-douyin/service/message/dal"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
)

func main() {
	dal.Init(context.Background())
	// client.Init(context.Background())
	r, err := etcd.NewEtcdRegistry([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constant.ServiceAddress+":"+constant.MessageServicePort)
	if err != nil {
		servLog.Fatal(err)
		return
	}
	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.MessageServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithTracer(
			prometheus.NewServerTracer(
				constant.MessageServerTracerPort,
				constant.MessageServerTracerPath)), // Tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	servLog.Warn("Message service started")
	err = svr.Run()

	if err != nil {
		servLog.Fatal(err)
	}
}
