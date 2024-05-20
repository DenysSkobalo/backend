package main

import (
	"auth_service/internal/api/routers"
	"auth_service/internal/database"
	"auth_service/internal/database/repositories"
	"auth_service/internal/server"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := server.LoadConfig()

	db, err := database.Connect(config.Database)
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	router := server.NewServer()

	repos := repositories.NewRepositories(db)

	routers.SetupRoutes(router, repos)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	// Handler to stop the server when a shutdown signal is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		logrus.Println("Shutting down server...")
		if err := srv.Shutdown(nil); err != nil {
			logrus.Error("Error shutting down server: ", err)
		}
	}()

	logrus.Printf("Server is running on port %s", config.Port)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Errorf("Error starting server: %v", err)
	}

}
