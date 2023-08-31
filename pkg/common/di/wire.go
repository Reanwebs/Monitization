package di

import (
	client "monit/pkg/client/conference"
	"monit/pkg/common/config"
	"monit/pkg/server"
)

func InitializeAPI(config config.Config) (*server.Server, error) {

	conferenceClient, err := client.InitClient(config)
	if err != nil {
		return nil, err
	}

	usecase := server.NewMonitizationServer(conferenceClient)

	server := server.NewGrpcServer(config, usecase)

	return server, nil
}
