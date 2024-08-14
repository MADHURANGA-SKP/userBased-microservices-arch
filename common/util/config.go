package util

import "github.com/spf13/viper"

type Config struct{
	JeagerAddr 		string `mapstructure:"JAEGER_ADDR"`
	HttpAddr 		string `mapstructure:"HTTP_SERVER_ADDR"`
	ConsulAddr 		string `mapstructure:"CONSUL_ADDR"`
	ServiceGateway  string `mapstructure:"SERVICE_GATEWAY"`
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