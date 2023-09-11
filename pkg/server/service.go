package server

import (
	"context"
	"log"
	"monit/pb/client/auth"
	pb "monit/pb/server"
	authClient "monit/pkg/client/auth"
	confClient "monit/pkg/client/conference"
	"monit/pkg/common/utils"
)

type monitizationServer struct {
	pb.UnimplementedMonitizationServer
	confClient confClient.ConferenceClient
	authClinet authClient.AuthClient
}

func NewMonitizationServer(confClient confClient.ConferenceClient, authClient authClient.AuthClient) *monitizationServer {
	return &monitizationServer{
		confClient: confClient,
		authClinet: authClient,
	}
}

func (m *monitizationServer) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Monitization server health checked")
	result := "Monitization server running"
	return &pb.Response{Result: result}, nil

}

func (m *monitizationServer) ParticipationReward(ctx context.Context, req *pb.ParticipationRewardRequest) (*pb.ParticipationRewardResponse, error) {
	coins, err := utils.CoinReward(req.Minutes, req.ConferenceType)
	if err != nil {
		return nil, err
	}

	authResp, err := m.authClinet.AddCoins(ctx, &auth.AddCoinsRequest{UserID: req.UserID, Coins: coins})
	if err != nil {
		return nil, err
	}

	resp := &pb.ParticipationRewardResponse{
		Result:    "Reward added succesfuly" + authResp.Result,
		CoinCount: coins,
	}
	return resp, nil
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
