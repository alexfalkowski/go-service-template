package main

import (
	"os"

	"github.com/alexfalkowski/go-service-template/cmd"
	scmd "github.com/alexfalkowski/go-service/cmd"
	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)
	command.AddVersion(cmd.Version)

	return command
}
