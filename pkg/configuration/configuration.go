package configuration

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/twrnakata/storefront-catalog-api/pkg/apperror"
)

type Environment struct {
	PORT              string
	DATABASE_URL      string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_DB       string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
}

var Env Environment

func InitConfig() error {
	viper.Reset()
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	Env = Environment{
		PORT:              strings.TrimSpace(viper.GetString("PORT")),
		DATABASE_URL:      strings.TrimSpace(viper.GetString("DATABASE_URL")),
		POSTGRES_HOST:     strings.TrimSpace(viper.GetString("POSTGRES_HOST")),
		POSTGRES_PORT:     strings.TrimSpace(viper.GetString("POSTGRES_PORT")),
		POSTGRES_DB:       strings.TrimSpace(viper.GetString("POSTGRES_DB")),
		POSTGRES_USER:     strings.TrimSpace(viper.GetString("POSTGRES_USER")),
		POSTGRES_PASSWORD: strings.TrimSpace(viper.GetString("POSTGRES_PASSWORD")),
	}

	if Env.PORT == "" {
		Env.PORT = "8080"
	}
	if Env.DATABASE_URL == "" {
		return apperror.ErrDatabaseURLRequired
	}

	return nil
}
