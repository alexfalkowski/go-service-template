package cmd

import "github.com/alexfalkowski/go-service/v2/cli"

// RegisterServer for cmd.
func RegisterServer(command cli.Commander) {
	cmd := command.AddServer("server", "Start go-service-template server", Module)

	cmd.AddInput("")
}
