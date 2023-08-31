package main

import (
	"log"
	"monit/pkg/common/config"
	"monit/pkg/common/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		if err := server.StartServer(config); err != nil {
			log.Fatalln(err)
		}

	}

}
