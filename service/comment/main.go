package main

import (
	"context"
	"net"
	"os"
	comment "simple-douyin/kitex_gen/comment/commentservice"
	"simple-douyin/pkg/constant"
	"simple-douyin/service/comment/client"
	"simple-douyin/service/comment/dal"

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

	addr, err := net.ResolveTCPAddr("tcp", constant.ServiceAddress+":"+constant.CommentServicePort)
	if err != nil {
		servLog.Fatal(err)
		return
	}
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.CommentServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithTracer(
			prometheus.NewServerTracer(
				constant.CommentServerTracerPort,
				constant.CommentServerTracerPath)), // Tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	servLog.Warn("Comment service started")
	err = svr.Run()

	if err != nil {
		servLog.Fatal(err)
	}
}
