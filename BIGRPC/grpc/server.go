package grpc

import (
	"log"
	"net"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
	"google.golang.org/grpc"
)

type Server struct {
	port   string
	server pb.NotificationServer
}

func NewGrpc(server pb.NotificationServer, port string) *Server {
	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) ListenAndServe() error {

	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, s.server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("failed to serve grpc server : ", err)
		return err
	}

	return nil

}
