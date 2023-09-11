package auth

import (
	"context"
	pb "monit/pb/client/auth"
)

type AuthClient interface {
	HealthCheck(context.Context, *pb.Request) (*pb.Response, error)
	AddCoins(context.Context, *pb.AddCoinsRequest) (*pb.AddCoinsResponse, error)
}
