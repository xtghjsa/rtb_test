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
	DSPHost        string `env:"DSP_HOST"`
	DSPPort        string `env:"DSP_PORT"`
	SSPHost        string `env:"sSP_HOST"`
	SSPPort        string `env:"ssp_PORT"`
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
	envList.DSPHost = os.Getenv("DSP_HOST")
	envList.DSPPort = os.Getenv("DSP_PORT")
	envList.SSPHost = os.Getenv("SSP_HOST")
	envList.SSPPort = os.Getenv("SSP_PORT")
	return envList, nil
}
