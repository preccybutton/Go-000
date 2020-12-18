package main

import (
	appinit "Go-000/Week04/cmd/init"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	pool, fn, err := appinit.InitStart("aaa")
	if err != nil{
		errors.WithMessage(err, "初始化失败")
		fmt.Printf("初始化失败  错误:\n%+v\n", err)
		return
	}

	exitChan := make(chan error)
	fmt.Println(pool.Server(), "---", pool.Lis())
	go func() {
		exitChan <- pool.Server().Serve(pool.Lis())
	}()
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	select {
	case <- quitChan:
		fn()
	case <- exitChan:
		fn()
	}
	return
}


