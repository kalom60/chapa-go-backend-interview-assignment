package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kalom60/chapa-go-backend-interview-assignment/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	fmt.Println("Here", cfg)
}
