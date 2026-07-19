package main

import (
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:              "8080",
		ReadTimeout:       15 * time.Second, // Потом парсить надо из env
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
