package server

import (
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
