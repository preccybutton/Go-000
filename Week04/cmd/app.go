package main

import "time"

type Option func(*options)

type options struct {
	closeTimeout  time.Duration
}

func (o *options) CloseTimeout() time.Duration{
	return o.closeTimeout
}


type App struct {
	opts *options
	cancel func()
	listen string
}

func (a *App) Opts() *options{
	return a.opts
}

func (a *App) Cancel() func(){
	return a.cancel
}

func (a *App) Listen() string{
	return a.listen
}

func SetCloseTimout(timeout time.Duration) Option{
	return func(i *options) {
		i.closeTimeout = timeout
	}
}


func New(lis string, opts ...Option) *App{
	options := new(options)
	for _, v := range opts{
		v(options)
	}
	return &App{opts:options, listen:lis}
}


func (a *App)Run() {

}