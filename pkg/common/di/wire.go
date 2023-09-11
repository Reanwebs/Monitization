package di

import (
	authClient "monit/pkg/client/auth"
	confClient "monit/pkg/client/conference"
	"monit/pkg/common/config"
	"monit/pkg/server"
)

func InitializeAPI(config config.Config) (*server.Server, error) {

	conferenceClient, err := confClient.InitConferenceClient(config)
	if err != nil {
		return nil, err
	}

	authClient, err := authClient.InitAuthClient(config)
	if err != nil {
		return nil, err
	}
	usecase := server.NewMonitizationServer(conferenceClient, authClient)

	server := server.NewGrpcServer(config, usecase)

	return server, nil
}
