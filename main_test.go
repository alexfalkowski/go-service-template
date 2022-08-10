//go:build features

package main

import (
	"testing"

	"github.com/alexfalkowski/go-service-template/cmd"
	scmd "github.com/alexfalkowski/go-service/cmd"
)

func TestFeatures(t *testing.T) {
	command := scmd.New()

	command.AddServer(cmd.ServerOptions)
	command.AddVersion(cmd.Version)

	if err := command.Run(); err != nil {
		t.Fatal(err.Error())
	}
}
