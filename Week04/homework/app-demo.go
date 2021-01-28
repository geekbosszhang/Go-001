package main

import (
	"context"
	"time"
	"os"
	"os/signal"
	"golang.org/x/sync/errgroup"
	"syscall"
)
type Option func(o *options)
type options struct{
	startTimeOut time.Duration
	stopTimeOut  time.Duration

	sigs []os.Signal
	sigFn func(*App, os.Signal)
}

type Hook struct{
	OnStart func(context.Context) error
	OnStop  func(context.Context) error
}
type App struct {
	opts options
	hooks []Hook 
	cancel func()
}
func New(opts ...Option) *App {
	options := options {
		startTimeOut: time.Second * 30,
		stopTimeOut: time.Second * 30,
		sigs: []os.Signal {
			syscall.SIGTERM,
			syscall.SIGQUIT,
			syscall.SIGINT,
		},
		sigFn: func(a *App, sig os.Signal) {
			switch sig {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				a.Stop()
			default:
			}
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return &App{opts: options}
}

func (a *App) Run() error {
	var ctx context.Context
	ctx, a.cancel = context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	for _, hook := range a.hooks {
		if hook.OnStop != nil {
			g.Go(func() error {
				<-ctx.Done()
				stopCtx, cancel := context.WithTimeout(context.Background(), a.opts.stopTimeOut)
				defer cancel()
				return hook.OnStop(stopCtx)
			})
		}
		if hook.OnStart != nil {
			g.Go(func() error {
				startContext, cancel := context.WithTimeout(context.Background(), a.opts.startTimeOut)
				defer cancel()
				return hook.OnStart(startContext)
			})
		}
	}
	if len(a.opts.sigs) > 0 {
		return g.Wait()
	}

	c := make(chan os.Signal, len(a.opts.sigs))
	signal.Notify(c, a.opts.sigs...)
	g.Go(func() error {
		for {
			select {
			case <- ctx.Done():
				return ctx.Err()
			case sig := <-c:
				if a.opts.sigFn != nil {
					a.opts.sigFn(a, sig)
				}
			}
		}
	})
	return g.Wait()
}

func (a *App) Stop() {
	if a.cancel != nil {
		a.cancel()
	}
}