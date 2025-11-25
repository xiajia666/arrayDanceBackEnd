package main

import (
	"arrayDanceBackEnd/douyin-api/globalinit/constant"
	"arrayDanceBackEnd/interaction-service/dao"
	"arrayDanceBackEnd/interaction-service/global"
	"arrayDanceBackEnd/interaction-service/handler"
	"arrayDanceBackEnd/interaction-service/proto/favorite"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	favorite.RegisterFavoriteServiceServer(server, &handler.FavoriteService{})

	listen, err := net.Listen("tcp", constant.FavoriteServiceAddr)
	if err != nil {
		panic(err)
	}

	dao.SetDefault(global.DB)

	global.ConsoleLogger.Info(constant.FavoriteServiceName,
		zap.String("Addr: ", constant.FavoriteServiceAddr),
	)
	global.InfoLogger.Info(constant.FavoriteServiceName,
		zap.String("Addr: ", constant.FavoriteServiceAddr),
	)

	err = server.Serve(listen)

	if err != nil {
		panic(err)
	}
}
