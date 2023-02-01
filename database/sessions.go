package database

import(
	"github.com/gofiber/storage/redis"
	"github.com/gofiber/fiber/v2/middleware/session"
)


var Store *session.Store

func ConnectRedis(){

	storage := redis.New(redis.Config{
		Host:      "redis",
		Port:      6379,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset: false,
	})
	
	Store = session.New(session.Config{
		Storage: storage,
	})
	
	
	}