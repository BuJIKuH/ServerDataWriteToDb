package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DataBaseInstance struct {
	DB *gorm.DB
}

var DataBase DataBaseInstance

func ConnectDataBase() error {
	conf := NewDataBase()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host, conf.Username, conf.Password, conf.DbName, conf.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n")
		os.Exit(2)
	}

	log.Println("Connect to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	//db.AutoMigrate(&ServerDbValue{})
	DataBase = DataBaseInstance{DB: db}
	return err
}
