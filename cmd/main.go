package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	application "user-service/internal/app"
	"user-service/internal/config"
	"user-service/internal/logger"
	"user-service/internal/storage"
)

func main() {
	fmt.Println("path is ", os.Getenv("CONFIG_PATH"))
	cfg := config.MustLoadConfig()
	db := storage.MustLoadDB(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	logger := logger.MustLoadLogger()

	app := application.NewApp(logger, db, cfg)

	go func() {
		app.MustRun()
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt, syscall.SIGTERM)
	<-q
	app.MustStop()
}
