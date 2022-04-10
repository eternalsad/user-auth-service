package app

import (
	"fmt"
	"net/http"

	"github.com/eternalsad/user-auth-service/config"
)

func Run() error {
	err := config.InitConfig()
	if err != nil {
		return fmt.Errorf("error during parsing config file: %s", err.Error())
	}
	repo := repository.NewRepository()
	service := service.NewService()
	handler := handler.NewHandler()

	srv := &http.Server{
		Handler: handler.InitRoutes(),
	}
	return nil
}
