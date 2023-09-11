package conference

import (
	"context"
	pb "monit/pb/client/conference"
	"monit/pkg/common/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type conferenceClient struct {
	Client pb.ConferenceClient
}

func InitConferenceClient(c config.Config) (ConferenceClient, error) {
	clientCon, err := grpc.Dial(c.ConfUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return NewConferenceClient(pb.NewConferenceClient(clientCon)), nil
}

func NewConferenceClient(client pb.ConferenceClient) ConferenceClient {
	return &conferenceClient{
		Client: client,
	}
}

func (c *conferenceClient) HealthCheck(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	resp, err := c.Client.HealthCheck(ctx, &pb.Request{
		Data: "Hi authentication server",
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
