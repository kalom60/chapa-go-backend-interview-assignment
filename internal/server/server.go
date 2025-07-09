package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/bank"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/transfer"
)

type Server struct {
	port     int
	bank     bank.Handler
	transfer transfer.Handler
}

func NewServer(port int, bank bank.Handler, transfer transfer.Handler) *http.Server {
	NewServer := &Server{
		port:     port,
		bank:     bank,
		transfer: transfer,
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
