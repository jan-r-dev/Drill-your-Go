package main

import (
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Error().
		Err(throwError()).
		Stack().
		Msg("Fibonacci is everywhere")

}

func throwError() error {
	return fmt.Errorf("unexpected error happened\nStack:\n%q", string(debug.Stack()))
}
