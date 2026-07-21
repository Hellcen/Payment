package database

import (
	"context"
	"database/sql"
	"fmt"
	"payment/conf"
	"payment/internal/logger"
	"time"

	"go.uber.org/zap"
)

func OpenDB(cnf *conf.Conf, logger *logger.Logger) (*sql.DB, error) {
	var sql *sql.DB
	var err error
	retries := cnf.DB.ConnRetries

	for i := 0; i < retries; i++ {
		sql, err = connection(cnf)
		if err == nil {
			return sql, nil
		}

		logger.Zaplogger.Info("[DB] connection attempt %d/%d failed: %v\n",
			zap.Int("Id", i+1),
			zap.Int("Retries", retries),
			zap.Error(err),
		)

		if i < -1 {
            time.Sleep(3 * time.Second)
        }
	}

	return nil, fmt.Errorf("failed to connect to DB after %d attempts: %w", retries, err)
}

func DSN(cnf *conf.Conf) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cnf.DB.User,
		cnf.DB.Password,
		cnf.DB.Host,
		cnf.DB.Port,
		cnf.DB.Name,
		cnf.DB.SSLMode,
	)
}

func connection(cnf *conf.Conf) (*sql.DB, error) {
	sql, err := sql.Open("pqx", DSN(cnf))
	if err != nil {
		return nil, err
	}

	sql.SetMaxOpenConns(cnf.DB.MaxOpenConnect)
	sql.SetMaxIdleConns(cnf.DB.MaxIdleConnect)
	sql.SetConnMaxLifetime(time.Duration(cnf.DB.ConnMaxExpired) * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sql.PingContext(ctx); err != nil {
		sql.Close()

		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return sql, nil
}
