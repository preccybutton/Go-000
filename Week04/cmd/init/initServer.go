package init

import (
	"Go-000/Week04/api"
	"Go-000/Week04/internal/server1/pkg"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
	"time"
)

type connPool struct {
	server *grpc.Server
	lis 	net.Listener
}

func (cp *connPool) Server() *grpc.Server{
	return cp.server
}

func (cp *connPool) Lis() net.Listener{
	return cp.lis
}

func InitServer(serc *serverConfig)(*connPool, func(), error){
	var err error
	lis, err := net.Listen("tcp", serc.lis)
	if err != nil{
		err = errors.Wrapf(err, "监听失败 地址: %v", serc.lis)
		return nil, nil, err
	}
	grpcServer := grpc.NewServer()
	api.RegisterGetTestDataServiceServer(grpcServer, pkg.InitService())
	fn := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		exit := make(chan struct{})
		go func(ch chan struct{}) {
			grpcServer.GracefulStop()
			exit <- struct{}{}
		}(exit)
		select {
		case <-exit:
			return
		case <- ctx.Done():
			grpcServer.Stop()
		}
		grpcServer.GracefulStop()
		grpcServer.Stop()
	}
	return &connPool{server:grpcServer, lis:lis}, fn, err
	//grpcServer.Serve()
	//grpcServer.GracefulStop()
	//grpcServer.Stop()
}