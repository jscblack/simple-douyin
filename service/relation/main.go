package main

import (
	"context"
	"net"
	"os"
	relation "simple-douyin/kitex_gen/relation/relationservice"
	"simple-douyin/pkg/constant"
	"simple-douyin/service/relation/client"
	"simple-douyin/service/relation/dal"

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
	client.Init(context.Background())
	r, err := etcd.NewEtcdRegistry([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constant.ServiceAddress+":"+constant.RelationServicePort)
	if err != nil {
		servLog.Fatal(err)
		return
	}
	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.RelationServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 1000}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithTracer(
			prometheus.NewServerTracer(
				constant.RelationServerTracerPort,
				constant.RelationServerTracerPath)), // Tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	servLog.Warn("Relation service started")
	err = svr.Run()

	if err != nil {
		servLog.Fatal(err)
	}
}
