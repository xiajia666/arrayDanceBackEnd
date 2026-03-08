package util

import (
	"demotest/douyin-api/global"
	"google.golang.org/grpc"
)

// 获取博客服务的gRPC连接
func GetBlogGrpcConn() (*grpc.ClientConn, error) {
	// 使用全局连接池,如果不存在则创建
	// 这里复用现有的连接管理逻辑
	// 临时方案:每次请求创建新连接
	// TODO: 优化为连接池
	// conn := global.GRPCBaseBlog
	// if conn == nil {
	// 	return nil, errors.New("博客服务连接未初始化")
	// }
	// return conn, nil

	// 临时方案
	// 如果需要长期方案,请在global/global.go中添加博客服务连接
	// 示例:
	// GRPCBaseBlog, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Blog gRPC server: %v", err)
	// }

	conn, err := grpc.Dial("127.0.0.1:8889", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
