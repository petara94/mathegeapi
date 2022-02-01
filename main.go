package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"mathegeapi/config"
	"mathegeapi/models"
)

func main() {
	cnf, err := config.LoadConfig(config.ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := gorm.Open(postgres.Open(cnf.Database.DSN()), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&models.Task{}, &models.TaskImage{})

	task := models.Task{}

	db.Preload(clause.Associations).First(&task)

	fmt.Println(cnf)

}
