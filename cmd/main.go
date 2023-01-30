package main

import (
	"coffeecatalogue4"
	"coffeecatalogue4/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(coffeecatalogue4.Server)
	if err := srv.Run("8080", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
