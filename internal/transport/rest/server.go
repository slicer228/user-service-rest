package rest

import (
	"context"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
	"time"
	um "user-service/internal/service/user-manager"
	"user-service/internal/transport/rest/middleware"
	"user-service/internal/transport/rest/users"
)

type HTTPServer struct {
	server *http.Server
	log    *slog.Logger
	host   string
	port   int
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
	r.Use(gin.Recovery())
	r.Use(middleware.RequestTimeoutMiddleware(requestTimeout))
	r.Use(requestid.New())

	//versions
	v1 := r.Group("/api/v1")
	{
		users.LoadUsersRouter(userManager, log, v1)
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

	return &server
}
