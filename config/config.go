package config

import (
	"os"

	"github.com/TamerB/ecommerce-stocks-tracker/constants"
)

type DBConfig struct {
	DBDriver string
	DBSource string
}

func NewConfig() *DBConfig {
	return &DBConfig{
		DBDriver: os.Getenv(constants.EnvDBDriver),
		DBSource: os.Getenv(constants.EnvDBSource),
	}
}
