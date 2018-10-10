// +build go1.8

package main

import (
	"os"

	"github.com/acquia/hub/commands"
	"github.com/acquia/hub/github"
	"github.com/acquia/hub/ui"
)

func main() {
	defer github.CaptureCrash()

	err := commands.CmdRunner.Execute()
	if !err.Ran {
		ui.Errorln(err.Error())
	}
	os.Exit(err.ExitCode)
}
