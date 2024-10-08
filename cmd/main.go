package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"users_service/internal/application"
	"users_service/internal/configuration"
	"users_service/internal/infrastructure/postgres"
	"users_service/internal/infrastructure/repository"
	"users_service/internal/infrastructure/sms"
	"users_service/internal/transport/http"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.New(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	smsService := sms.New(&config.SMS)

	repo := repository.New(db)
	useCase := application.New(repo, smsService, config.FS.UploadFolder)

	httpServer := http.New(useCase, config.FS.UploadFolder)
	httpServer.Register()
	go httpServer.ListenAndServe(config.Server.HttpSocket)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdown()

	httpServer.Shutdown(ctx)
	db.Close()
}
