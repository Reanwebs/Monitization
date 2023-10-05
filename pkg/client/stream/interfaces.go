package stream

import (
	"context"
	pb "monit/pb/client/stream"
)

type VideoServiceClient interface {
	HealthCheck(context.Context, *pb.Request) (*pb.Response, error)
	VideoDetails(context.Context, *pb.VideoDetailsRequest) (*pb.VideoDetailsResponse, error)
}
