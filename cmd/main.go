package main

import (
	"hotel/internal/handler"
	"hotel/internal/repositories"
	"hotel/internal/services"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	storage := repositories.NewStorage()
	service := services.NewServices(storage)
	handler := handler.NewHandler(service)

	server := &http.Server{
		Addr:    ":5050",
		Handler: handler.InitRoutes(),
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}

}
