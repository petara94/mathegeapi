package main

import (
	"mathegeapi/config"
	"mathegeapi/database"
)

func init() {
	config.Setup(config.ConfigFilePath)
	database.Setup()
}

func main() {

}
