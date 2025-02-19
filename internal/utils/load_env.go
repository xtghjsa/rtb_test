package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// List of environment variables
type EnvList struct {
	PostgresUser   string `env:"POSTGRES_USER"`
	PostgresPass   string `env:"POSTGRES_PASSWORD"`
	PostgresDBName string `env:"POSTGRES_DBNAME"`
	PostgresHost   string `env:"POSTGRES_HOST"`
	PostgresPort   string `env:"POSTGRES_PORT"`
	ServerHost     string `env:"SERVER_HOST"`
	ServerPort     string `env:"SERVER_PORT"`
}

// LoadEnv loads environment variables and returns them as a struct
func LoadEnv() (EnvList, error) {
	envList := EnvList{}
	err := godotenv.Load()
	if err != nil {
		return envList, err
	}
	envList.PostgresUser = os.Getenv("POSTGRES_USER")
	envList.PostgresPass = os.Getenv("POSTGRES_PASSWORD")
	envList.PostgresDBName = os.Getenv("POSTGRES_DBNAME")
	envList.PostgresHost = os.Getenv("POSTGRES_HOST")
	envList.PostgresPort = os.Getenv("POSTGRES_PORT")
	envList.ServerHost = os.Getenv("SERVER_HOST")
	envList.ServerPort = os.Getenv("SERVER_PORT")
	return envList, nil
}
