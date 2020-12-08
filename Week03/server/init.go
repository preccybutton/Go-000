package server

import (
	"context"
	"errors"
	"net/http"
	"time"
)

var Stopped bool

func StartServe(ctx context.Context, addr string, handler http.HandlerFunc, isExist chan <- error, stop <- chan struct{}) error {
	s := http.Server{
		Addr: addr,
		Handler: handler,
	}
	go func() {
		select {
		case <- ctx.Done():
			ctx2, _  := context.WithTimeout(context.Background(), 2*time.Second)
			s.Shutdown(ctx2)
			isExist <- nil

		case <- stop:
			isExist <- s.Shutdown(context.Background())
		}
	}()
	return s.ListenAndServe()
}
