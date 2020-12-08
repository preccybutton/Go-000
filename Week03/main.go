package main

import (
	myserve "Go-000/Week03/server"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)


func main(){
	done := make(chan error, 1)
	isExist := make(chan error, 2)
	stop := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(
		func() error {
			return myserve.StartServe(ctx,":19001", myserve.FirstHandlec, isExist, stop)
		})
	g.Go(
		func() error {
			return myserve.StartServe(ctx,":19002", myserve.SecondHandlec, isExist, stop)
		})
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	//程序自己异常退出的信号传递
	go func() {
		done <- g.Wait()
	}()
	select {
	//user主动关闭
	case <- quitChan:
		close(stop)
		fmt.Println("shutdown by signal")
		//程序自己异常
	case err := <- done :
		fmt.Printf("error is %v", err)
	}
	for i := 0; i < 2; i ++{
		<- isExist
	}
	return
}
