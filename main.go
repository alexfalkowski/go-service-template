package main

import (
	"github.com/alexfalkowski/go-service-template/internal/cmd"
	"github.com/alexfalkowski/go-service/v2/cli"
	"github.com/alexfalkowski/go-service/v2/context"
)

var app = cli.NewApplication(func(command cli.Commander) {
	cmd.RegisterServer(command)
})

func main() {
	app.ExitOnError(context.Background())
}
