package main

import (
	"context"
	"net"
	"os"
	publish "simple-douyin/kitex_gen/publish/publishservice"
	"simple-douyin/pkg/constant"
	"simple-douyin/service/publish/client"
	"simple-douyin/service/publish/dal"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("RUN_MODE") == "Production" {
		servLog.SetLevel(servLog.WarnLevel)
	}
	// init db
	dal.Init(context.Background())
	// init rpc client
	client.Init(context.Background())
	r, err := etcd.NewEtcdRegistry([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constant.ServiceAddress+":"+constant.PublishServicePort)
	if err != nil {
		servLog.Fatal(err)
		return
	}
	var svr server.Server
	if os.Getenv("BENCHMARK_MODE") == "True" {
		svr = publish.NewServer(new(PublishServiceImpl),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.PublishServiceName}), // server name
			// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
			// server.WithMiddleware(middleware.ServerMiddleware),
			server.WithServiceAddr(addr), // address
			// server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 1000}), // limit
			server.WithMuxTransport(), // Multiplex
			server.WithTracer(
				prometheus.NewServerTracer(
					constant.PublishServerTracerPort,
					constant.PublishServerTracerPath)), // Tracer
			// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
			server.WithRegistry(r), // registry
		)
	} else {
		svr = publish.NewServer(new(PublishServiceImpl),
			server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.PublishServiceName}), // server name
			// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
			// server.WithMiddleware(middleware.ServerMiddleware),
			server.WithServiceAddr(addr),                                         // address
			server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 1000}), // limit
			server.WithMuxTransport(),                                            // Multiplex
			server.WithTracer(
				prometheus.NewServerTracer(
					constant.PublishServerTracerPort,
					constant.PublishServerTracerPath)), // Tracer
			// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
			server.WithRegistry(r), // registry
		)
	}
	servLog.Warn("Publish service started")
	err = svr.Run()

	if err != nil {
		servLog.Fatal(err)
	}
}
