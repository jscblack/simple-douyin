package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/prometheus/common/log"
	"net"
	publish "simple-douyin/kitex_gen/publish/publishservice"
	"simple-douyin/pkg/constant"
	"simple-douyin/service/publish/client"
	"simple-douyin/service/publish/dal"
)

func main() {
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
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.PublishServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		// server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	err = svr.Run()

	if err != nil {
		servLog.Fatal(err)
	}
}
