package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	DBHost       string `mapstructure:"DB_HOST"`
	UserName     string `mapstructure:"DB_USER"`
	UserPassword string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBPort       int    `mapstructure:"DB_PORT"`
}

type RedisConfig struct {
	RedisHost    string `mapstructure:"REDIS_Host"`
	UserName     string `mapstructure:"REDIS_Username"`
	UserPassword string `mapstructure:"REDIS_Password"`
	RedisPort    int    `mapstructure:"REDIS_Port"`
}

func LoadDBConfig(path string) (config DBConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadRedisConfig(path string) (config RedisConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
