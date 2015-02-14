package info

import (
	"github.com/codegangsta/cli"
)


func InfoCommands() []cli.Command {
	c := []cli.Command {
		ShowRevisionCountCommand(),
	}
	return c
}

