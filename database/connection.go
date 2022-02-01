package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mathegeapi/config"
	"mathegeapi/models"
)

var DB *gorm.DB

func Setup() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Config.Database.DSN()), &gorm.Config{})

	if err != nil {
		log.Fatalf("database.Setup() err: %v", err)
	}

	_ = DB.AutoMigrate(&models.Task{}, &models.TaskImage{}, &models.PatternTask{})
}
