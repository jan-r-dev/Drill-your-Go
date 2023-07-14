package main

import (
	"flag"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	logFile, err := os.Create("./logs/output.log")
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := initLogging(logFile)

	if err := doWork(logger); err != nil {
		logger.Error().
			Stack().
			Err(err).
			Msg("")
	}

}

func initLogging(logFile io.Writer) zerolog.Logger {
	zerolog.SetGlobalLevel(parseSeverityFlag())
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return zerolog.New(logFile).With().Timestamp().Logger()
}

func parseSeverityFlag() zerolog.Level {
	var severity string
	flag.StringVar(&severity, "severity", "info", "Configure logging severity. Available levels are:\n\terror\n\tinfo\n\tdebug\n")

	flag.Parse()

	switch severity {
	case "error":
		return zerolog.ErrorLevel
	case "debug":
		return zerolog.DebugLevel
	default:
		return zerolog.InfoLevel
	}
}

func postRequisitesWithError() error {
	return errors.New("unexpected error happened")
}

func doWork(logger zerolog.Logger) error {
	logger.Info().
		Str("event", "task finished").
		Msg("")

	logger.Debug().
		Int("tasks remaining", 1).
		Str("event", "post requisites started").
		Msg("")

	err := postRequisitesWithError()

	logger.Debug().
		Int("tasks remaining", 0).
		Str("event", "post requisites finished with error").
		Msg("")

	return err
}
