package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"payment/conf"
	"payment/internal/logger"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	cnf := conf.NewConf()
	cnf = conf.Parse(cnf)

	logger, err := logger.NewLogger()
	if err != nil {
		panic(fmt.Sprintf("Failed to init logger: %v", err))
	}
	defer logger.Zaplogger.Sync()

	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Server health")
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
		logger.Zaplogger.Info("Server started",
			zap.String("port", cnf.Server.Addr),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Zaplogger.Fatal("Server failed start",
				zap.Error(err),
			)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
	)
	<-q

	logger.Zaplogger.Info("Server shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Zaplogger.Fatal("Server shutdown failed",
			zap.Error(err),
		)
	}

	logger.Zaplogger.Info("Server Closed")
}
