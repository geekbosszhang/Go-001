package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serverAPP(ctx context.Context) error {
	s := &http.Server{
		Addr:    ":8080",
		Handler: &DemoHandler{},
	}
	go func() {
		<-ctx.Done()
		s.Shutdown(ctx)
		fmt.Println("server has shut down")
	}()
	return s.ListenAndServe()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			//stop server
		}()
		// start server
		return serverAPP(ctx)
	})

	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)

		for {
			fmt.Println("signal")
			select {
			case sign := <-sig:
				fmt.Println(sign)
				return nil
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			}
		}
	})

}
