package main

import (
	"os"

	"github.com/alexfalkowski/go-service-template/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "env:CONFIG_FILE")
	c.AddServer("server", "Start go-service-template server", cmd.ServerOptions...)

	return c
}
