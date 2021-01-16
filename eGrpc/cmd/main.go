package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aivuca/goms/eGrpc/internal/server/grpc"
	"github.com/aivuca/goms/eGrpc/internal/server/http"
)

func main() {
	fmt.Println("\n---eGrpc---")

	httpSrv := http.New()
	log.Printf("new http server: %+v", httpSrv)

	grpcSrv := grpc.New()
	log.Printf("new grpc server: %+v", grpcSrv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
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
