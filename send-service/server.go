package main

import (
	"fmt"
	"log"
	"net"
	"path"
	"runtime"
	"textgrpc/configs"
	"textgrpc/send"

	"google.golang.org/grpc"
)

const (
	port = 9002
)

func main() {
	// 加载配置文件
	_, file, _, _ := runtime.Caller(0)
	configPath := path.Dir(file) + "/configs/config.json"
	configs.LoadConfig(configPath)

	rpcServer := grpc.NewServer()
	send.RegisterSendServiceServer(rpcServer, new(send.SendService))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("start send rpc: %s", lis.Addr())
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
