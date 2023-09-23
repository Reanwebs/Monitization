package main

import (
	"log"
	"monit/pkg/common/config"
	"monit/pkg/common/di"
	"monit/pkg/common/utils"
	"monit/pkg/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := server.ConnectPsqlDB(config, &utils.Wallet{}, &utils.UserRewardHistory{})
	if err != nil {
		log.Fatalln(err)
	}
	server, err := di.InitializeAPI(config, db)
	if err != nil {
		log.Fatalln(err)
	} else {
		if err := server.StartServer(config); err != nil {
			log.Fatalln(err)
		}
	}
}
