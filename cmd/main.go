package main

import (
	"coffeecatalogue4"
	"coffeecatalogue4/pkg/handler"
	"coffeecatalogue4/pkg/logging"
	"coffeecatalogue4/pkg/repository"
	"coffeecatalogue4/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(coffeecatalogue4.Server)
	if err := srv.Run(viper.GetString("path"), handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
