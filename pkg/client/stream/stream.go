package stream

import (
	"context"
	pb "monit/pb/client/stream"
	"monit/pkg/common/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type videoServiceClient struct {
	Client pb.VideoServiceClient
}

func InitVideoClient(c config.Config) (VideoServiceClient, error) {
	clientCon, err := grpc.Dial(c.StreamUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewConferenceClient(pb.NewVideoServiceClient(clientCon)), nil
}

func NewConferenceClient(client pb.VideoServiceClient) VideoServiceClient {
	return &videoServiceClient{
		Client: client,
	}
}

func (c *videoServiceClient) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	resp, err := c.Client.HealthCheck(ctx, &pb.Request{
		Data: "Hi authentication server",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *videoServiceClient) VideoDetails(ctx context.Context, req *pb.VideoDetailsRequest) (*pb.VideoDetailsResponse, error) {

	resp, err := c.Client.VideoDetails(ctx, &pb.VideoDetailsRequest{VideoID: req.VideoID})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
