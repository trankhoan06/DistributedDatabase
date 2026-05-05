package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var config *Configuration

type Configuration struct {
	ThaiLanXml         string `env:"THAILANXML"`
	ThaiLanReplicaXml  string `env:"THAILANREPLICAXML"`
	VietNamXml         string `env:"VIETNAMXML"`
	VietNamReplicaXml  string `env:"VIETNAMREPLICAXML"`
	CambodiaXml        string `env:"CAMBODIAXML"`
	CambodiaReplicaXml string `env:"CAMBODIAREPLICAXML"`
}

func GetConfig() *Configuration {
	if config == nil {
		err := godotenv.Load()
		fmt.Println(err)
		config = &Configuration{}
		if err := env.Parse(config); err != nil {
			log.Fatal().Err(err).Msg("Get Config Error")
		}
	}
	return config
}
