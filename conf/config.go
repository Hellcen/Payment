package conf

import "time"

type Conf struct {
	Server ServerConf
	DB     DBConfig
}

type ServerConf struct {
	Addr              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

type DBConfig struct {
}
