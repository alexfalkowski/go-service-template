package main

import (
	"context"

	"github.com/alexfalkowski/go-service-template/internal/cmd"
	sc "github.com/alexfalkowski/go-service/cmd"
)

var app = sc.NewApplication(func(command *sc.Command) {
	cmd.RegisterServer(command)
})

func main() {
	app.ExitOnError(context.Background())
}
