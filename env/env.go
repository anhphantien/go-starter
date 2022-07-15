package env

import (
	"strconv"

	"github.com/spf13/viper"
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
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	PORT = viper.GetString("PORT")

	DB_USER = viper.GetString("DB_USER")
	DB_PASS = viper.GetString("DB_PASS")
	DB_HOST = viper.GetString("DB_HOST")
	DB_PORT = viper.GetString("DB_PORT")
	DB_NAME = viper.GetString("DB_NAME")

	JWT_EXPIRES_AT, _ = strconv.Atoi(viper.GetString("JWT_EXPIRES_AT"))
	JWT_SECRET = []byte(viper.GetString("JWT_SECRET"))
}
