package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	addr string
	echo *echo.Echo
	db   *sqlx.DB
}

func NewServer(addr string, db *sqlx.DB) *Server {
	return &Server{
		addr: addr,
		echo: echo.New(),
		db:   db,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr: s.addr,
	}

	go func() {
		if err := s.echo.StartServer(server); err != nil {
			log.Fatalf("Could not start the server: %v", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		log.Fatalf("An error has occurred mapping the handlers: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)

	<-quit

	log.Println("Server is stopping..")

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.echo.Shutdown(ctx)
}
