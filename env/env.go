package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT string

	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string

	JWT_EXPIRES_AT int
	JWT_SECRET     []byte
)

func init() {
	godotenv.Load(".env")

	PORT = os.Getenv("PORT")

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")

	JWT_EXPIRES_AT, _ = strconv.Atoi(os.Getenv("JWT_EXPIRES_AT"))
	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
}
