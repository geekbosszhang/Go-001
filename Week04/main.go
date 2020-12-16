package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"net/http"
	"fmt"
	"os"
	"syscall"
	"os/signal"
	"time"
)

func startServer(server *http.Server) error{
	e := InitializeEvent()
    e.Start()
	return server.ListenAndServe()
}

func main()  {
	// create a context that we can cancel
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	defer cancel()

	// create a new mux and handler
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("server: received request")
		time.Sleep(3 * time.Second)
		fmt.Println("server: request finished")
	}))

	// create a server
	srv := &http.Server{Addr: ":8080", Handler: mux}

	g.Go(func() error {
		go func() {
			<- ctx.Done()
			fmt.Println("http ctx done")
			// shut down gracefully, but wait no longer than 5 seconds before halting
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			srv.Shutdown(shutdownCtx)
			fmt.Println("server shut down")
		}()
		return startServer(srv)
	})
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("recived signal")
			select {
			case <- ctx.Done():
				fmt.Println("signal ctx done")
				return ctx.Err()
			case <- sig:
				cancel()
				return nil
			}
		}
	})

	err := g.Wait()
	fmt.Println(err)

}