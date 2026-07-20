package conf

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

func Parse(cnf *Conf) *Conf {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: File .env not found")
	}

	err = cleanenv.ReadEnv(cnf)
	if err != nil {
		log.Fatalf("File .env failed to parse: %v", err)
	}

	return cnf
}