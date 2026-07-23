package conf

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Conf struct {
	Server ServerConf
	DB     DBConfig
}

type ServerConf struct {
	Addr              string        `env:"ADDR" env-default:":8080"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT" env-default:"20s"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" env-default:"20s"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT" env-default:"30s"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT" env-default:"120s"`
}

type DBConfig struct {
	Host     string `env:"HOST" env-default:"localhost"`
	Port     int    `env:"PORT" env-default:"5432"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"PASSWORD" env-default:"postgres"`
	Name     string `env:"NAME" env-default:"postgres"`
	SSLMode  string `env:"SSLMODE" env-default:"disable"`

	MaxIdleConnect int           `env:"MAX_IDLE_CONNECTION" env-default:"20"`
	MaxOpenConnect int           `env:"MAX_OPEN_CONNECTION" env-default:"50"`
	ConnMaxExpired time.Duration `env:"CONN_MAX_EXPIRED" env-default:"5m"`
	ConnRetries    int           `env:"CONNECT_RETRIES" env-default:"15"`
}

func NewConf() (*Conf, error) {
	var cnf Conf

	return Parse(&cnf)
}

// interface zapcore.ObjectMarshaler
func (s *ServerConf) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ADDR", s.Addr)
	enc.AddDuration("READ_TIMEOUT", s.ReadTimeout)
	enc.AddDuration("READ_HEADER_TIMEOUT", s.ReadHeaderTimeout)
	enc.AddDuration("WRITE_TIMEOUT", s.WriteTimeout)
	enc.AddDuration("IDLE_TIMEOUT", s.IdleTimeout)

	return nil
}

// interface zapcore.ObjectMarshaler
func (d *DBConfig) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("HOST", d.Host)
	enc.AddInt("PORT", d.Port)
	enc.AddString("USER", d.User)
	enc.AddString("NAME", d.Name)
	enc.AddString("SSLMODE", d.SSLMode)

	enc.AddInt("MAX_IDLE_CONNECTION", d.MaxIdleConnect)
	enc.AddInt("MAX_OPEN_CONNECTION", d.MaxOpenConnect)
	enc.AddDuration("CONN_MAX_EXPIRED", d.ConnMaxExpired)
	enc.AddInt("CONNECT_RETRIES", d.ConnRetries)

	return nil
}

// interface zapcore.ObjectMarshaler
func (c *Conf) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddObject("Server", &c.Server)
	enc.AddObject("Database", &c.DB)

	return nil
}
