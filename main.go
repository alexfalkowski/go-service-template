package main

import (
	"github.com/alexfalkowski/go-service-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	command().ExitOnError()
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:CONFIG_FILE")
	c.AddServer("server", "Start go-service-template server", cmd.ServerOptions...)

	return c
}
