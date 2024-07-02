package logger

import (
	"os"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger = log.Logger

func init() {
	err := envy.Load("../../.env")
	if err != nil {

	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logLevel := envy.Get("LOG_LEVEL", "info")
	l, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Warn().Err(err).Msg("failed to parse log level, using info level")
	}

	zerolog.SetGlobalLevel(l)

	Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}
