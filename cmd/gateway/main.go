package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"payment/conf"
	"syscall"
	"time"
)

func main() {
	cnf := conf.NewConf()
	cnf = conf.Parse(cnf)

	fmt.Println(cnf)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w,"Server health")
	})

	server := &http.Server{
		Addr:              cnf.Server.Addr,
		Handler:           mux,
		ReadTimeout:       cnf.Server.ReadTimeout,
		ReadHeaderTimeout: cnf.Server.ReadHeaderTimeout,
		WriteTimeout:      cnf.Server.WriteTimeout,
		IdleTimeout:       cnf.Server.IdleTimeout,
	}

	go func() {
		log.Printf("Server started: Port - %s", cnf.Server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
	)
	<-q

	log.Println("Server shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server closed")
}
