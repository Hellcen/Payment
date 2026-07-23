package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"payment/RegistrationService/conf"
	"payment/RegistrationService/internal/adapter"
	"payment/RegistrationService/internal/database"
	"payment/RegistrationService/internal/handler"
	"payment/RegistrationService/internal/repository/postgreSQL"
	"payment/RegistrationService/internal/service"
	"payment/pkg/logger"

	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	cnf, err := conf.NewConf()
	logger, err := logger.NewLogger()
	goValidator := adapter.New()

	if err != nil {
		panic(fmt.Sprintf("Failed to init logger: %v", err))
	}
	defer logger.Zaplogger.Sync()

	logger.Zaplogger.Info("Config",
		zap.Object("cnf", cnf),
	)

	db, err := database.OpenDB(cnf, logger)
	if err != nil {
		logger.Zaplogger.Error("DB init failed",
			zap.Error(err),
		)
		logger.Zaplogger.Fatal("DB no connection")
	}
	defer db.Close()

	logger.Zaplogger.Info("Database connected",
		zap.Int("port", cnf.DB.Port),
	)

	postgresRep := postgresql.NewRepository(db)
	userRep := service.NewAuthService(postgresRep)
	authHandler := handler.NewAuthHandler(userRep, *logger, goValidator)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "Server health")
	})

	mux.HandleFunc("POST /api/v1/register", authHandler.RegisterHandler)

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
	defer signal.Stop(q)
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
