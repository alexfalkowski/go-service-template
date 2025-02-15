package main

import (
	"github.com/alexfalkowski/go-service-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/env"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(env.NewVersion().String())
	command.RegisterInput(command.Root(), "env:CONFIG_FILE")

	cmd.RegisterServer(command)

	return command
}
