package main

import (
	"context"
	"net"
	relation "simple-douyin/kitex_gen/relation/relationservice"
	"simple-douyin/pkg/constant"
	"simple-douyin/service/user/dal"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/prometheus/common/log"
)

func main() {
	// init db
	dal.Init(context.Background())
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
