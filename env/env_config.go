package env

import (
	"log"
	"os"
)

var (
	SECRET_KEY string
	PORT       string
	DB_URL     string
)

func init() {
	SECRET_KEY = os.Getenv("SECRET_KEY")

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	DB_URL = os.Getenv("DB_URL")
	if DB_URL == "" {
		DB_URL = "dev.db"
	}

	log.Println("Variáveis carregadas")
}
