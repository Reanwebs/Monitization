package server

import (
	"context"
	"log"
	pb "monit/pb/server"
	client "monit/pkg/client/conference"
)

type monitizationServer struct {
	pb.UnimplementedMonitizationServer
	confClient client.ConferenceClient
}

func NewMonitizationServer(confClient client.ConferenceClient) *monitizationServer {
	return &monitizationServer{
		confClient: confClient,
	}
}

func (m *monitizationServer) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Monitization server health checked")
	result := "Monitization server running"
	return &pb.Response{Result: result}, nil

}
