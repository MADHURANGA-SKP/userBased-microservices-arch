package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct{
	JeagerAddr 				string `mapstructure:"JAEGER_ADDR"`
	HttpAddr 				string `mapstructure:"HTTP_SERVER_ADDR"`
	GRPCAddr 				string `mapstructure:"GRPC_ADDR"`
	RMQUser 				string `mapstructure:"RABBITMQ_USER"`
	RMQPass 				string `mapstructure:"RABBITMQ_PASS"`
	RMQHost 				string `mapstructure:"RABBITMQ_HOST"`
	RMQPort 				string `mapstructure:"RABBITMQ_PORT"`
	ConsulAddr 				string `mapstructure:"CONSUL_ADDR"`
	DBSorurce 				string `mapstructure:"DB_SOURCE"`
	ServiceGateway  		string `mapstructure:"SERVICE_GATEWAY"`
	ServiceUser  			string `mapstructure:"SERVICE_USER"`
	MigrationURL  			string `mapstructure:"MIGRATION_URL"`
	TokenSymmetricKey 		string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  	time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration 	time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string)(config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return 
	}

	err = viper.Unmarshal(&config)
	return
}