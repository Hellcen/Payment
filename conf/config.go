package conf

import (
	"time"
)

type Conf struct {
	Server ServerConf
	DB     DBConfig
}

type ServerConf struct {
	Addr              string        `env:"ADDR"`
	ReadTimeout       time.Duration `env:"READ_TIMEOUT"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT"`
	WriteTimeout      time.Duration `env:"WRITE_TIMEOUT"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT"`
}

type DBConfig struct {
}

func NewConf() *Conf {
	return &Conf{
		Server: ServerConf{
			Addr:              "8080",
			ReadTimeout:       30 * time.Second,
			ReadHeaderTimeout: 20 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       180 * time.Second,
		},
	}
}
