package server

import (
	"log"
	pb "monit/pb/server"
	"monit/pkg/common/config"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	GrpcServer   *grpc.Server
	PortListener net.Listener
}

func NewGrpcServer(cfg config.Config, server pb.MonitizationServer) *Server {
	listener, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterMonitizationServer(grpcServer, server)
	return &Server{GrpcServer: grpcServer, PortListener: listener}
}

func (s *Server) StartServer(cfg config.Config) error {
	log.Println("Monitization service is running...", cfg.Port)
	return s.GrpcServer.Serve(s.PortListener)
}
