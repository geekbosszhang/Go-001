package main
import (
"context"
"errors"
"fmt"
"golang.org/x/sync/errgroup"
"net/http"
"os"
"os/signal"
"syscall"
)

type DemoHandler struct {

}

func (h *DemoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func serverApp(ctx context.Context)  error{
	s:=&http.Server{
		Addr:":8080",
		Handler: &DemoHandler{},
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
		fmt.Println("serverApp has shutdown")
	}()

	fmt.Println("server begin to start...")
	return s.ListenAndServe()
}

func listenSignal(ctx context.Context) error {
	c:=make(chan os.Signal,1)
	signal.Notify(c,syscall.SIGQUIT,syscall.SIGTERM,syscall.SIGINT)
	select {
	case <-c:
		fmt.Println("signal is", c)
	case <-ctx.Done():
		fmt.Println("receive ctx done in listenSignal")
	}
	signal.Stop(c)
	close(c)
	return errors.New("quit")
}

func main() {
	g,ctx:=errgroup.WithContext(context.Background())

	g.Go(func() error {
		return serverApp(ctx)
	})

	g.Go(func() error {
		return listenSignal(ctx)
	})

	if err:=g.Wait(); err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("main done")
}