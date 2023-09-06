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

func (m *monitizationServer) UserRewardHistory(ctx context.Context, req *pb.UserRewardHistoryRequest) (*pb.UserRewardHistoryResponse, error) {
	return nil, nil
}

func (m *monitizationServer) GroupRewardHistory(ctx context.Context, req *pb.GroupRewardHistoryRequest) (*pb.GroupRewardHistoryResponse, error) {
	return nil, nil
}

func (m *monitizationServer) UserWatchHour(ctx context.Context, req *pb.UserWatchHourRequest) (*pb.UserWatchHourResponse, error) {
	return nil, nil
}

func (m *monitizationServer) GroupWatchHour(ctx context.Context, req *pb.GroupWatchHourRequest) (*pb.GroupWatchHourResponse, error) {
	return nil, nil
}

func (m *monitizationServer) ConferenceWatchHour(ctx context.Context, req *pb.ConferenceWatchHourRequest) (*pb.ConferenceWatchHourResponse, error) {
	return nil, nil
}
