package database

import (
	"fmt"
	"log"
	"os"

	"github.com/AhmedFawzy0/TO-DO/app/models"
	"github.com/AhmedFawzy0/TO-DO/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {

	config, err := config.LoadDBConfig(".")
	if err != nil {
		panic("Cannot load config")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Canada/Eastern",
		config.DBHost,
		config.UserName,
		config.UserPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.User{}, &models.Task{})

	DB = Dbinstance{

		Db: db,
	}

}
