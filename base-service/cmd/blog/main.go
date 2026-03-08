package main

import (
	"demotest/base-service/handler/blog"
	bpb "demotest/base-service/proto/blog"
	"flag"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 8889, "端口号")

	// 添加blog.proto
	flag.Parse()
	fmt.Print("ip: ", *IP)
	fmt.Print("  port: ", *Port)
	fmt.Println("  Blog Service is running")

	server := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second, // 服务器在收到 keep-alive ping 后的最小等待时间
			PermitWithoutStream: true,            // 允许在没有活动流的情况下发送 keep-alive ping
		}),
	)

	// 注册博客服务
	bpb.RegisterBlogServiceServer(server, &blog.BlogServe{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen " + err.Error())
	}

	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
}
