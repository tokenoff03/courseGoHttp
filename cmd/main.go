package main

import (
	"hotel/internal/handler"
	"hotel/internal/repositories"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()

	storage := repositories.NewStorage()
	handler := handler.NewHandler(storage)

	mux.HandleFunc("/room/", handler.RoomHandler)
	mux.HandleFunc("/room", handler.RoomHandler)

	server := &http.Server{
		Addr:    ":5050",
		Handler: mux,
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}

}
