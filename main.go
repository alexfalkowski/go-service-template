package main

import (
	"os"

	"github.com/alexfalkowski/go-service-template/cmd"
	scmd "github.com/alexfalkowski/go-service/cmd"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

func command() *scmd.Command {
	command := scmd.New(cmd.Version)

	command.AddServer(cmd.ServerOptions...)

	return command
}
