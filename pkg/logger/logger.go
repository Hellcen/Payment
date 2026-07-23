package logger

import "go.uber.org/zap"

type Logger struct {
	Zaplogger *zap.Logger
}

func NewLogger() (*Logger, error){
	zapLogger, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	return &Logger{Zaplogger: zapLogger}, nil
}

func(l *Logger) Sync() error {
	return l.Zaplogger.Sync()
}