package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse: ", log.LstdFlags)

	srv := server.New(logger)

	logger.Println("server started on :8080")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
