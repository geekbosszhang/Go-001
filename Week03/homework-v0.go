//基于errgroup实现一个http server的启动和关闭，以及Linux signal的信号的注册和处理，要保证一个退出，全部注销退出
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
	"net/http"
	"golang.org/x/sync/errgroup"
	"io"
)

func main() {
	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)
	defer cancel()

	// a WaitGroup for the goroutines to tell us they've stopped
	wg := sync.WaitGroup{}

	// a channel for `server()` to tell us they've stopped
	wg.Add(1)
	g.Go(func() error {
		// tell the caller that we've stopped
		defer wg.Done()

		// create a new mux and handler
		mux := http.NewServeMux()
		mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("server: received request")
			time.Sleep(3 * time.Second)
			io.WriteString(w, "Finished!\n")
			fmt.Println("server: request finished")
		}))

		// create a server
		srv := &http.Server{Addr: ":8080", Handler: mux}

		go func() {
			// service connections
			if err := srv.ListenAndServe(); err != nil {
				fmt.Printf("Listen : %s\n", err)
			}
		}()

		<-gctx.Done()
		fmt.Println("server: caller has told us to stop")

		// shut down gracefully, but wait no longer than 5 seconds before halting
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// ignore error since it will be "Err shutting down server : context canceled"
		srv.Shutdown(shutdownCtx)

		fmt.Println("server gracefully stopped")
		return nil
	})

	// listen for C-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("main: received Ctrl-c - shutting down")

	// tell the goroutines to stop
	fmt.Println("main: telling goroutines to stop")
	cancel()

	// and wait for them to reply back
	wg.Wait()
	fmt.Println("main: all goroutines have told us they've finished")
}
