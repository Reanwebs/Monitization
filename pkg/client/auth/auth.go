package auth

import (
	"context"
	pb "monit/pb/client/auth"
	"monit/pkg/common/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	Client pb.AuthClient
}

func InitAuthClient(c config.Config) (AuthClient, error) {
	clientCon, err := grpc.Dial(c.ConfUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewConferenceClient(pb.NewAuthClient(clientCon)), nil
}

func NewConferenceClient(client pb.AuthClient) AuthClient {
	return &authClient{
		Client: client,
	}
}

func (c *authClient) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	resp, err := c.Client.HealthCheck(ctx, &pb.Request{
		Data: "Hi authentication server",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *authClient) AddCoins(ctx context.Context, req *pb.AddCoinsRequest) (*pb.AddCoinsResponse, error) {
	return nil, nil
}
