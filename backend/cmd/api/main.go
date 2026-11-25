package main

import (
	"flag"
	"github.com/kyawzinthant-coding/froky/internal/modules"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
	"strings"
	"time"
)

var (
	logLevel = flag.String("log-level", "debug", "Log level (debug, info, warn, error, fatal, panic)")
)

func main() {
	flag.Parse()
	setupLogger()

	log.Info().Msg("Starting app...")

	app := fx.New(
		modules.Config,
		modules.Database,
		modules.AuthModule,
		modules.MenuModule,
		modules.VendorModule,
		modules.HandlerModule,
		modules.ServerModule,
	)

	app.Run()
}

func setupLogger() {
	level := parseLogLevel(*logLevel)
	zerolog.SetGlobalLevel(level)

	zerolog.TimeFieldFormat = time.RFC3339

	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).
		With().
		Timestamp().
		Caller().
		Logger()

	log.Debug().Str("level", *logLevel).Msg("Logger initialized")
}

func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}
