package middleware

import (
	"github.com/AhmedFawzy0/TO-DO/config"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var Store *session.Store

func ConnectRedis() {

	config, err := config.LoadRedisConfig(".")
	if err != nil {
		panic("Cannot load config")
	}

	storage := redis.New(redis.Config{
		Host:     config.RedisHost,
		Port:     config.RedisPort,
		Username: config.UserName,
		Password: config.UserPassword,
		Database: 0,
		Reset:    false,
	})

	Store = session.New(session.Config{
		Storage: storage,
	})

}
