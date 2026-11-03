package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Warn().
			Err(err).
			Msg(".env file not found, using environment variables")
	} else {
		log.Debug().Msg(".env file loaded")
	}

	cfg := &AppConfig{
		DataBase:   &DatabaseConfig{},
		CorsConfig: &CorsConfig{},
		Http:       &HttpConfig{},
	}

	opts := env.Options{
		RequiredIfNoDef: true,
	}

	log.Debug().Msg("Parsing environment variables into config")
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to load configuration: missing or invalid env vars")
	}

	log.Info().
		Str("db_host", cfg.DataBase.Host).
		Str("db_name", cfg.DataBase.Name).
		Int("http_port", cfg.Http.Port).
		Msg("Configuration loaded")

	return cfg

}
