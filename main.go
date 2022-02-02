package main

import (
	"log"
	"mathegeapi/config"
	"mathegeapi/routers"
	"mathegeapi/stores"
	"net/http"
	"os"
)

func init() {
}

func main() {
	cnf, err := config.LoadConfig(config.ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.MkdirAll(cnf.Server.ImagesDir, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}

	store := stores.NewStore(&cnf.Database)
	err = store.Open()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := routers.GetRouter(cnf, store)

	err = http.ListenAndServe(":"+cnf.Server.Port, r)
	if err != nil {
		log.Fatal(err)
	}
}
