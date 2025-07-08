package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/bank"
)

type Server struct {
	port int
	bank bank.Handler
}

func NewServer(port int, bank bank.Handler) *http.Server {
	NewServer := &Server{
		port: port,
		bank: bank,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
