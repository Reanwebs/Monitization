package conference

import (
	"context"
	pb "monit/pb/client/conference"
)

type ConferenceClient interface {
	HealthCheck(context.Context, *pb.Request) (*pb.Response, error)
}
