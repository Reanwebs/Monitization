package server

import (
	"context"
	"log"
	"monit/pb/client/auth"
	pb "monit/pb/server"
	authClient "monit/pkg/client/auth"
	confClient "monit/pkg/client/conference"
	"monit/pkg/common/utils"
	"monit/pkg/repository"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type monitizationServer struct {
	pb.UnimplementedMonitizationServer
	confClient confClient.ConferenceClient
	authClinet authClient.AuthClient
	userRepo   repository.UserRepoMethods
}

func NewMonitizationServer(confClient confClient.ConferenceClient, authClient authClient.AuthClient, repo repository.UserRepoMethods) *monitizationServer {
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

func (m *monitizationServer) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.CreateWalletResponse, error) {
	if err := m.userRepo.CreateWallet(req.UserID); err != nil {
		return nil, err
	}
	return &pb.CreateWalletResponse{Result: "Wallet Created"}, nil
}

func (m *monitizationServer) GetWallet(ctx context.Context, req *pb.GetWalletRequest) (*pb.GetWalletResponse, error) {
	response, err := m.userRepo.GetWallet(req.UserID)
	if err != nil {
		return nil, err
	}
	return &pb.GetWalletResponse{Coins: int32(response.Coins)}, nil
}

func (m *monitizationServer) UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.UpdateWalletResponse, error) {
	var rewardType string
	if req.Reason == "Referal" {
		rewardType = "Credit"
	} else {
		rewardType = req.Type
	}
	input := utils.UserRewardHistory{
		UserID:          req.UserID,
		RewardReason:    req.Reason,
		TransactionType: rewardType,
		Referal:         req.UserName,
	}
	if err := m.userRepo.UpdateWalletHistory(input); err != nil {
		return nil, err
	}
	if err := m.userRepo.UpdateWallet(req.UserID, 10); err != nil {
		return nil, err
	}
	return &pb.UpdateWalletResponse{Result: "Wallet Updated"}, nil
}

func (m *monitizationServer) UserRewardHistory(ctx context.Context, req *pb.UserRewardHistoryRequest) (*pb.UserRewardHistoryResponse, error) {
	var result *[]utils.UserRewardHistory
	var err error
	if req.Sort == "" {
		result, err = m.userRepo.GetRewardHistory(req.UserID)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = m.userRepo.GetRewardHistoryBySort(req.UserID, req.Sort)
		if err != nil {
			return nil, err
		}
	}
	response := &pb.UserRewardHistoryResponse{}
	for _, item := range *result {
		pbItem := &pb.UserRewardHistory{
			UserID:          item.UserID,
			RewardReason:    item.RewardReason,
			TransactionType: item.TransactionType,
			CoinCount:       uint32(item.CoinCount),
			Time:            timestamppb.New(item.Time),
		}
		response.Result = append(response.Result, pbItem)
	}
	return response, nil
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
