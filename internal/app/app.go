package app

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/atadzan/audio-library/internal/config"
	v1 "github.com/atadzan/audio-library/internal/pkg/handler/v1"
	"github.com/atadzan/audio-library/internal/pkg/repository"
	"github.com/atadzan/audio-library/internal/pkg/repository/postgres"
	"github.com/atadzan/audio-library/internal/pkg/storage"
	"github.com/atadzan/audio-library/internal/pkg/storage/minIO"
	"github.com/gofiber/fiber/v2"
)

func startServerWithGracefulShutdown(app *fiber.App, port string) {
	idleConnClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnClosed)
	}()

	// Run server.
	if err := app.Listen(":" + port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnClosed
}

func Init(configPath string) error {
	ctx := context.Background()
	appConfig, err := config.Load(configPath)
	if err != nil {
		return err
	}
	dbPool, err := postgres.NewPoolConn(ctx, appConfig.App.DatabaseURL)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	repo := repository.New(dbPool)
	minioClient, err := minIO.NewMinioClient(appConfig.Storage)
	if err != nil {
		return err
	}
	newStorage := storage.New(minioClient)

	handlers := v1.NewHandler(repo, newStorage)

	app := fiber.New(fiber.Config{
		StrictRouting:     true,
		EnablePrintRoutes: true,
	})

	handlers.InitRoutes(app)

	startServerWithGracefulShutdown(app, appConfig.App.Port)

	return nil
}
