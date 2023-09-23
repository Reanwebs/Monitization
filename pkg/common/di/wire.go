package di

import (
	authClient "monit/pkg/client/auth"
	confClient "monit/pkg/client/conference"
	"monit/pkg/common/config"
	"monit/pkg/repository"
	"monit/pkg/server"

	"gorm.io/gorm"
)

func InitializeAPI(config config.Config, db *gorm.DB) (*server.Server, error) {
	conferenceClient, err := confClient.InitConferenceClient(config)
	if err != nil {
		return nil, err
	}
	authClient, err := authClient.InitAuthClient(config)
	if err != nil {
		return nil, err
	}
	userRepo := repository.NewUserRepo(db)
	usecase := server.NewMonitizationServer(conferenceClient, authClient, userRepo)
	server := server.NewGrpcServer(config, usecase)
	return server, nil
}
