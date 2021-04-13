package server

import (
	"log"
	"net"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
	"google.golang.org/grpc"
)

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()

	pb.RegisterNotificationServer(s, &NotificationServer{})

	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
