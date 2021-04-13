package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/arijitnayak92/go-lang-projects/BIGRPC/domain"
	grpcPkg "github.com/arijitnayak92/go-lang-projects/BIGRPC/grpc"
	"github.com/arijitnayak92/go-lang-projects/BIGRPC/grpchandler"
)

var (

	// grpc server port
	grpcServerPort string

	// program controller
	done    = make(chan struct{})
	errGrpc = make(chan error)
)

func init() {

	flag.StringVar(&grpcServerPort, "grpcServerPort", ":5060", "grpc server port")

}

func handleInterrupts() {
	log.Println("start handle interrupts")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	sig := <-interrupt
	log.Printf("caught sig: %v", sig)
	// close resource here
	done <- struct{}{}
}

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go handleInterrupts()

	d := domain.NewAMSNotification()

	grpcHandler := grpchandler.NewGrpcHandler(d)
	grpcServer := grpcPkg.NewGrpc(grpcHandler, grpcServerPort)

	go func() {
		fmt.Printf("GRPC sever running on port: %v\n", grpcServerPort)
		errGrpc <- grpcServer.ListenAndServe()
	}()

	select {
	case err := <-errGrpc:
		log.Print("Grpc error", err)
	case <-done:
		log.Println("shutting down server ...")
	}
	time.AfterFunc(1*time.Second, func() {
		close(done)
		close(errGrpc)
	})
}
