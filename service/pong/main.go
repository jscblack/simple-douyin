package main

import (
	"net"
	"simple-douyin/kitex_gen/pong/pongservice"

	"simple-douyin/pkg/constant"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	servLog "github.com/sirupsen/logrus"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constant.EtcdAddressWithPort})
	if err != nil {
		servLog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constant.ServiceAddress+":"+constant.PingServicePort)
	if err != nil {
		servLog.Fatal(err)
		return
	}
	svr := pongservice.NewServer(new(PongServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.PingServiceName}), // server name
		// server.WithMiddleware(middleware.CommonMiddleware),                                            // middleWare
		// server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 1000}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		// server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r), // registry
	)
	servLog.Warn("Ping service started")
	err = svr.Run()

	if err != nil {
		servLog.Fatal(err)
	}
}
