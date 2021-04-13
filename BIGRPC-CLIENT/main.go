package main

import (
	"log"

	"github.com/arijitnayak92/go-lang-projects/BIGRPC-CLIENT/service"
	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
	"google.golang.org/grpc"
)

func main() {

	// dail server
	conn, err := grpc.Dial(":5060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	laptopClient := pb.NewNotificationClient(conn)
	notificationClient := service.NewNotificationClient(laptopClient)
	s := service.NewService(notificationClient)

	s.NotificationService.PushNotification()
}
