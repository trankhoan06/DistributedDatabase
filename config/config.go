package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var config *Configuration

type Configuration struct {
	VersionApp string `env:"VERSION_APP"`
	PortApp    string `env:"PORT_APP"`
	ModeApp    string `env:"MODE_APP"`
	PrefixApp  string `env:"PREFIX_APP"`
	SecretApp  string `env:"SECRET_APP"`
	HostApp    string `env:"HOST_APP"`
	//VerifyKey    *rsa.PublicKey
	//SignKey      *rsa.PrivateKey
	MigrationURLApp      string `env:"MIGRATION_URL_APP"`
	HostMysql            string `env:"HOST_MYSQL"`
	ContainerNameMysql   string `env:"CONTAINER_MYSQL"`
	PortMysql            string `env:"PORT_MYSQL"`
	UserMysql            string `env:"USER_MYSQL"`
	PasswordMysql        string `env:"PASSWORD_MYSQL"`
	DBNameMysql          string `env:"DB_MYSQL"`
	MaxOpenConnsMysql    int    `env:"MAX_OPEN_CONNS_MYSQL"`
	MaxIdleConnsMysql    int    `env:"MAX_IDLE_CONNS_MYSQL"`
	MaxConnLifetimeMysql string `json:"maxConnLifetimeMysql"`
	EmailSenderName      string `env:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string `env:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string `env:"EMAIL_SENDER_PASSWORD"`
	ResendApiKey         string `env:"RESEND_API_KEY"`
	HostRedis            string `env:"HOST_REDIS"`
	PortRedis            string `env:"PORT_REDIS"`
	PasswordRedis        string `env:"PASSWORD_REDIS"`
	DBNameRedis          string `env:"DB_REDIS"`
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
