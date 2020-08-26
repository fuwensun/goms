package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aivuca/goms/eHttp/internal/server/http"
)

func main() {
	fmt.Println("\n---eHttp---")

	httpSrv := http.New()
	log.Printf("new http server: %+v", httpSrv)

	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sch
		log.Printf("get a signal %v", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("server exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
