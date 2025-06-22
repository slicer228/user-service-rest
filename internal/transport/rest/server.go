package rest

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
	"net/http"
	"strings"
	"time"
	_ "user-service/docs"
	um "user-service/internal/service/user-manager"
	"user-service/internal/transport/rest/handlers/users"
	"user-service/internal/transport/rest/middleware"
)

// @title           User Service API
// @version         1.0
// @description     REST API для обогащения пользовательским данных
// @Host            localhost:8080
// @BasePath        /api/v1
// @schemes         http

type HTTPServer struct {
	server  *http.Server
	log     *slog.Logger
	Address string
}

func (s *HTTPServer) MustStart() {
	err := s.server.ListenAndServe()
	if err != nil {
		s.log.Error(err.Error())
		panic(err)
	}
}

func (s *HTTPServer) GracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)

	if err != nil {
		s.log.Error("Error when stopping server", "err", err.Error())
		panic(err)
	}
}

func NewHTTPServer(log *slog.Logger, userManager *um.UserManager, requestTimeout *time.Duration, host string, port string) *HTTPServer {
	var server HTTPServer

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gin.Recovery())
	r.Use(middleware.RequestTimeoutMiddleware(requestTimeout))
	r.Use(requestid.New())

	//versions
	v1 := r.Group("/api/v1")
	{
		usersHandler := users.NewUsersRouter(log, userManager, v1)
		usersHandler.Load()
	}

	var address strings.Builder

	address.WriteString(host)
	address.WriteString(":")
	address.WriteString(port)

	httpServer := &http.Server{
		Addr:    address.String(),
		Handler: r,
	}

	server.server = httpServer
	server.log = log
	server.Address = address.String()

	return &server
}
