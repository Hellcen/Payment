package conf

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

func Parse(cnf *Conf) (*Conf, error) {
	_ = godotenv.Load()

	if err := cleanenv.ReadEnv(cnf); err != nil {
		return nil, fmt.Errorf("File .env failed to parse: %v", err)
	}

	return cnf, nil
}