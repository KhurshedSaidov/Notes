package main

import (
	"log"
	"net/http"
	"notes/handlers"
	"notes/internal/configs"
)

func main() {
	Run()
}

func Run() error {
	config, err := configs.InitConfigs()
	if err != nil {
		return err
	}
	router := handlers.InitRouters()
	address := config.Ip + config.Port

	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Println("listen and serve error", err)
		return err
	}
	return nil
}
