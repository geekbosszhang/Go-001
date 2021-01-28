package cmd

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	flag.Parse()
	log.Println("start")
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <- c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Println("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
