package main

import (
	"log"
	"net"
	pong "simple-douyin/kitex_gen/pong/pongservice"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		log.Println(err.Error())
		return
	}
	svr := pong.NewServer(new(PongServiceImpl), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
