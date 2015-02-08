package work

import (
	"github.com/codegangsta/cli"
)

func WorkCommands() []cli.Command {
	workCommands := []cli.Command {
		ShowCatalogCommand(),
		ShowTasksCommand(),
	}
	return workCommands
}

