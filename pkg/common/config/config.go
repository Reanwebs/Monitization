package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	AuthUrl    string `mapstructure:"AUTH_SERVER_URL"`
	ConfUrl    string `mapstructure:"CONFERENCE_SERVER_URL"`
	StreamUrl  string `mapstructure:"STREAM_SERVER_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
