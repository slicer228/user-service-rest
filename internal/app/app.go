package app

import (
	"log/slog"
	"user-service/internal/config"
	um "user-service/internal/service/user-manager"
	"user-service/internal/storage"
	"user-service/internal/transport/rest"
)

type App struct {
	HTTPServer *rest.HTTPServer
	log        *slog.Logger
}

func (a *App) MustRun() {
	a.log.Debug("Starting app...")
	a.log.Debug("Starting server...")
	a.HTTPServer.MustStart()
}

func (a *App) MustStop() {
	a.log.Debug("Shutting down...")
	a.HTTPServer.GracefulShutdown()
	a.log.Debug("Shut down complete")
}

func NewApp(log *slog.Logger, db *storage.DBConnection, cfg *config.Config) *App {
	log.Debug("Initializing user manager service...")
	userManager := um.NewUserManager(db, log, cfg)
	return &App{
		HTTPServer: rest.NewHTTPServer(log, userManager, &cfg.Timeout, cfg.Host, cfg.Port),
		log:        log,
	}
}
