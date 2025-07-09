package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kalom60/chapa-go-backend-interview-assignment/config"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/bank"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/cache"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/repository"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/server"
	"github.com/kalom60/chapa-go-backend-interview-assignment/internal/transfer"

	chapa "github.com/Chapa-Et/chapa-go"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, closeDB, err := repository.OpenDB(cfg.DbUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer closeDB()

	redis, err := cache.NewRedis(cfg.RedisUrl, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err.Error())
	}

	store := repository.NewStore(db)
	bank := bank.New(store)
	chapaAPI := chapa.New()
	transfer := transfer.New(cfg.WebhookSecret, store, chapaAPI, redis)

	srv := server.NewServer(cfg.Port, bank, transfer)

	shutdownError := make(chan error)

	if cfg.Render != "true" {
		go func() {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit

			log.Println("shutting down server...")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			shutdownError <- srv.Shutdown(ctx)
		}()
	}

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}

	if err := <-shutdownError; err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
