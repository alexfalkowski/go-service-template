package main

import (
	"github.com/alexfalkowski/go-service-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	command := sc.New(cmd.Version)
	command.RegisterInput(command.Root(), "env:CONFIG_FILE")

	cmd.RegisterServer(command)

	return command
}
