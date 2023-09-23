package server

import (
	"fmt"
	"log"
	"monit/pkg/common/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPsqlDB(cfg config.Config, models ...interface{}) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DbHost, cfg.DbUser, cfg.DbName, cfg.DbPort, cfg.DbPassword)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	db.AutoMigrate(models...)
	return db, nil
}
